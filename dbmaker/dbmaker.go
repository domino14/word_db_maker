// Package dbmaker creates SQLITE databases for various lexica, so I can use
// them in my word game empire.
package dbmaker

import (
	"bufio"
	"database/sql"
	"fmt"
	"github.com/domino14/macondo/gaddag"
	"github.com/domino14/macondo/lexicon"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"sort"
	"strings"
)

type Alphagram struct {
	words        []string
	combinations uint64
	alphagram    string
}

func (a *Alphagram) String() string {
	return fmt.Sprintf("Alphagram: %s (%d)", a.alphagram, a.combinations)
}

type AlphByCombos []Alphagram // used to be []*Alphagram

func (a AlphByCombos) Len() int      { return len(a) }
func (a AlphByCombos) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a AlphByCombos) Less(i, j int) bool {
	// XXX: Existing aerolith dbs don't sort by alphagram to break ties.
	// (It's sort of random unfortunately)
	// The DBs generated by this tool will be slightly off. We must continue
	// to use the old DBs until there is a lexicon update :(
	if a[i].combinations == a[j].combinations {
		return a[i].alphagram < a[j].alphagram
	} else {
		return a[i].combinations > a[j].combinations
	}
}

type LexiconMap map[string]lexicon.LexiconInfo

type LexiconSymbolDefinition struct {
	In     string // The word is in this lexicon
	NotIn  string // The word is not in this lexicon
	Symbol string // The corresponding lexicon symbol
}

// create a sqlite db for this lexicon name.
func createSqliteDb(lexiconName string) string {
	dbName := "./" + lexiconName + ".db"
	os.Remove(dbName)
	sqlStmt := `
	CREATE TABLE alphagrams (probability int, alphagram varchar(20),
	    length int, combinations int);

	CREATE TABLE words (word varchar(20), alphagram varchar(20),
	    lexicon_symbols varchar(5), definition varchar(512),
	    front_hooks varchar(26), back_hooks varchar(26),
	    inner_front_hook int, inner_back_hook int);

	CREATE INDEX alpha_index on alphagrams(alphagram);
	CREATE INDEX prob_index on alphagrams(probability, length);
	CREATE INDEX word_index on words(word);
	CREATE INDEX alphagram_index on words(alphagram);
	`
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}
	return dbName
}

func CreateLexiconDatabase(lexiconName string, lexiconInfo lexicon.LexiconInfo,
	lexSymbols []LexiconSymbolDefinition, lexMap LexiconMap) {
	fmt.Println("Creating lexicon database", lexiconName)
	definitions, alphagrams := populateAlphsDefs(lexiconInfo.LexiconFilename,
		lexiconInfo.Combinations, lexiconInfo.LetterDistribution)
	fmt.Println("Sorting by probability")
	alphs := alphaMapValues(alphagrams)
	sort.Sort(AlphByCombos(alphs))

	var probs [16]uint32
	for i := 0; i < 16; i++ {
		probs[i] = 0
	}

	dbName := createSqliteDb(lexiconName)

	alphInsertQuery := `
	INSERT INTO alphagrams(probability, alphagram, length, combinations)
	VALUES (?, ?, ?, ?)`
	wordInsertQuery := `
	INSERT INTO words (word, alphagram, lexicon_symbols, definition,
		front_hooks, back_hooks, inner_front_hook, inner_back_hook)
	VALUES(?, ?, ?, ?, ?, ?, ?, ?)`

	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatal(err)
	}
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	alphStmt, err := tx.Prepare(alphInsertQuery)
	if err != nil {
		log.Fatal(err)
	}
	wordStmt, err := tx.Prepare(wordInsertQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer alphStmt.Close()
	defer wordStmt.Close()
	gd := lexiconInfo.Gaddag
	for idx, alph := range alphs {
		if idx%10000 == 0 {
			log.Println(idx, "...")
		}
		wl := len([]rune(alph.alphagram))
		if wl <= 15 {
			probs[wl]++
		}
		_, err = alphStmt.Exec(probs[wl], alph.alphagram, wl, alph.combinations)
		if err != nil {
			log.Fatal(err)
		}
		for _, word := range alph.words {
			if err != nil {
				log.Fatal(err)
			}

			backHooks := sortedHooks(gaddag.FindHooks(gd, word, gaddag.BackHooks),
				lexiconInfo.LetterDistribution)
			frontHooks := sortedHooks(gaddag.FindHooks(gd, word, gaddag.FrontHooks),
				lexiconInfo.LetterDistribution)
			frontInnerHook := 0
			backInnerHook := 0
			if gaddag.FindInnerHook(gd, word, gaddag.BackInnerHook) {
				backInnerHook = 1
			}
			if gaddag.FindInnerHook(gd, word, gaddag.FrontInnerHook) {
				frontInnerHook = 1
			}

			def := definitions[word]
			alphagram := alph.alphagram

			wordStmt.Exec(
				word, alphagram,
				findLexSymbols(word, lexiconName, lexMap, lexSymbols), def,
				frontHooks, backHooks, frontInnerHook, backInnerHook)
		}
	}
	tx.Commit()
}

func sortedHooks(hooks []rune, dist lexicon.LetterDistribution) string {
	w := lexicon.Word{Word: string(hooks), Dist: dist}
	return w.MakeAlphagram()
}

func findLexSymbols(word string, lexiconName string, lexMap LexiconMap,
	lexSymbols []LexiconSymbolDefinition) string {

	symbols := ""

	for _, def := range lexSymbols {
		if lexiconName == def.In {
			lex := lexMap[def.NotIn]
			if !gaddag.FindWord(lex.Gaddag, word) {
				symbols += def.Symbol
			}
		}
	}
	return symbols
}

// The values of the map.
func alphaMapValues(theMap map[string]Alphagram) []Alphagram {
	x := make([]Alphagram, len(theMap))
	i := 0
	for _, value := range theMap {
		x[i] = value
		i++
	}
	return x
}

func populateAlphsDefs(filename string, combinations func(string) uint64,
	dist lexicon.LetterDistribution) (
	map[string]string, map[string]Alphagram) {
	definitions := make(map[string]string)
	alphagrams := make(map[string]Alphagram)
	file, _ := os.Open(filename)
	// XXX: Check error
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) > 0 {
			word := lexicon.Word{Word: strings.ToUpper(fields[0]), Dist: dist}
			definition := ""
			if len(fields) > 1 {
				definition = strings.Join(fields[1:], " ")
			}
			definitions[word.Word] = definition
			alphagram := word.MakeAlphagram()
			alph, ok := alphagrams[alphagram]
			if !ok {
				alphagrams[alphagram] = Alphagram{
					[]string{word.Word},
					combinations(alphagram),
					alphagram}
			} else {
				alph.words = append(alph.words, word.Word)
				alphagrams[alphagram] = alph
			}
		}
	}
	file.Close()
	return definitions, alphagrams
}
