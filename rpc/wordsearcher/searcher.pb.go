// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc/wordsearcher/searcher.proto

package wordsearcher

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SearchRequest_Condition int32

const (
	SearchRequest_LEXICON            SearchRequest_Condition = 0
	SearchRequest_LENGTH             SearchRequest_Condition = 1
	SearchRequest_PROBABILITY_RANGE  SearchRequest_Condition = 2
	SearchRequest_PROBABILITY_LIST   SearchRequest_Condition = 3
	SearchRequest_PROBABILITY_LIMIT  SearchRequest_Condition = 4
	SearchRequest_NUMBER_OF_ANAGRAMS SearchRequest_Condition = 5
	SearchRequest_NUMBER_OF_VOWELS   SearchRequest_Condition = 6
	// HAS_TAGS = 7;
	SearchRequest_POINT_VALUE      SearchRequest_Condition = 8
	SearchRequest_MATCHING_ANAGRAM SearchRequest_Condition = 9
	SearchRequest_ALPHAGRAM_LIST   SearchRequest_Condition = 10
	SearchRequest_NOT_IN_LEXICON   SearchRequest_Condition = 11
	SearchRequest_REGEX            SearchRequest_Condition = 12
	// A list of words.
	SearchRequest_WORD_LIST SearchRequest_Condition = 13
)

var SearchRequest_Condition_name = map[int32]string{
	0:  "LEXICON",
	1:  "LENGTH",
	2:  "PROBABILITY_RANGE",
	3:  "PROBABILITY_LIST",
	4:  "PROBABILITY_LIMIT",
	5:  "NUMBER_OF_ANAGRAMS",
	6:  "NUMBER_OF_VOWELS",
	8:  "POINT_VALUE",
	9:  "MATCHING_ANAGRAM",
	10: "ALPHAGRAM_LIST",
	11: "NOT_IN_LEXICON",
	12: "REGEX",
	13: "WORD_LIST",
}

var SearchRequest_Condition_value = map[string]int32{
	"LEXICON":            0,
	"LENGTH":             1,
	"PROBABILITY_RANGE":  2,
	"PROBABILITY_LIST":   3,
	"PROBABILITY_LIMIT":  4,
	"NUMBER_OF_ANAGRAMS": 5,
	"NUMBER_OF_VOWELS":   6,
	"POINT_VALUE":        8,
	"MATCHING_ANAGRAM":   9,
	"ALPHAGRAM_LIST":     10,
	"NOT_IN_LEXICON":     11,
	"REGEX":              12,
	"WORD_LIST":          13,
}

func (x SearchRequest_Condition) String() string {
	return proto.EnumName(SearchRequest_Condition_name, int32(x))
}

func (SearchRequest_Condition) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_14226a2334c91a98, []int{2, 0}
}

type SearchRequest_NotInLexCondition int32

const (
	SearchRequest_OTHER_ENGLISH    SearchRequest_NotInLexCondition = 0
	SearchRequest_PREVIOUS_VERSION SearchRequest_NotInLexCondition = 1
)

var SearchRequest_NotInLexCondition_name = map[int32]string{
	0: "OTHER_ENGLISH",
	1: "PREVIOUS_VERSION",
}

var SearchRequest_NotInLexCondition_value = map[string]int32{
	"OTHER_ENGLISH":    0,
	"PREVIOUS_VERSION": 1,
}

func (x SearchRequest_NotInLexCondition) String() string {
	return proto.EnumName(SearchRequest_NotInLexCondition_name, int32(x))
}

func (SearchRequest_NotInLexCondition) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_14226a2334c91a98, []int{2, 1}
}

// An Alphagram encapsulates info about an alphagram, including the words,
// length, probability, combinations.
type Alphagram struct {
	Alphagram string  `protobuf:"bytes,1,opt,name=alphagram,proto3" json:"alphagram,omitempty"`
	Words     []*Word `protobuf:"bytes,2,rep,name=words,proto3" json:"words,omitempty"`
	// expandedRepr is true if the length, probability, combinations are
	// included. Otherwise, this is an "unexpanded" alphagram.
	// Note that if expandedRepr is true, then the `words` field is also
	// expanded (with definition, hooks, etc).
	ExpandedRepr         bool     `protobuf:"varint,3,opt,name=expandedRepr,proto3" json:"expandedRepr,omitempty"`
	Length               int32    `protobuf:"varint,4,opt,name=length,proto3" json:"length,omitempty"`
	Probability          int32    `protobuf:"varint,5,opt,name=probability,proto3" json:"probability,omitempty"`
	Combinations         int64    `protobuf:"varint,6,opt,name=combinations,proto3" json:"combinations,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Alphagram) Reset()         { *m = Alphagram{} }
func (m *Alphagram) String() string { return proto.CompactTextString(m) }
func (*Alphagram) ProtoMessage()    {}
func (*Alphagram) Descriptor() ([]byte, []int) {
	return fileDescriptor_14226a2334c91a98, []int{0}
}

func (m *Alphagram) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Alphagram.Unmarshal(m, b)
}
func (m *Alphagram) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Alphagram.Marshal(b, m, deterministic)
}
func (m *Alphagram) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Alphagram.Merge(m, src)
}
func (m *Alphagram) XXX_Size() int {
	return xxx_messageInfo_Alphagram.Size(m)
}
func (m *Alphagram) XXX_DiscardUnknown() {
	xxx_messageInfo_Alphagram.DiscardUnknown(m)
}

var xxx_messageInfo_Alphagram proto.InternalMessageInfo

func (m *Alphagram) GetAlphagram() string {
	if m != nil {
		return m.Alphagram
	}
	return ""
}

func (m *Alphagram) GetWords() []*Word {
	if m != nil {
		return m.Words
	}
	return nil
}

func (m *Alphagram) GetExpandedRepr() bool {
	if m != nil {
		return m.ExpandedRepr
	}
	return false
}

func (m *Alphagram) GetLength() int32 {
	if m != nil {
		return m.Length
	}
	return 0
}

func (m *Alphagram) GetProbability() int32 {
	if m != nil {
		return m.Probability
	}
	return 0
}

func (m *Alphagram) GetCombinations() int64 {
	if m != nil {
		return m.Combinations
	}
	return 0
}

// A Word is more than just the string representing the word. It has other
// info like the definition, hooks, lex symbols, etc.
type Word struct {
	Word      string `protobuf:"bytes,1,opt,name=word,proto3" json:"word,omitempty"`
	Alphagram string `protobuf:"bytes,2,opt,name=alphagram,proto3" json:"alphagram,omitempty"`
	// Note: the following fields are not filled in if the alphagram's
	// `expandedRepr` is false. Protobuf fields are optional already, but
	// this lets us be explicit.
	Definition           string   `protobuf:"bytes,3,opt,name=definition,proto3" json:"definition,omitempty"`
	FrontHooks           string   `protobuf:"bytes,4,opt,name=front_hooks,json=frontHooks,proto3" json:"front_hooks,omitempty"`
	BackHooks            string   `protobuf:"bytes,5,opt,name=back_hooks,json=backHooks,proto3" json:"back_hooks,omitempty"`
	LexiconSymbols       string   `protobuf:"bytes,6,opt,name=lexicon_symbols,json=lexiconSymbols,proto3" json:"lexicon_symbols,omitempty"`
	InnerFrontHook       bool     `protobuf:"varint,7,opt,name=inner_front_hook,json=innerFrontHook,proto3" json:"inner_front_hook,omitempty"`
	InnerBackHook        bool     `protobuf:"varint,8,opt,name=inner_back_hook,json=innerBackHook,proto3" json:"inner_back_hook,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Word) Reset()         { *m = Word{} }
func (m *Word) String() string { return proto.CompactTextString(m) }
func (*Word) ProtoMessage()    {}
func (*Word) Descriptor() ([]byte, []int) {
	return fileDescriptor_14226a2334c91a98, []int{1}
}

func (m *Word) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Word.Unmarshal(m, b)
}
func (m *Word) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Word.Marshal(b, m, deterministic)
}
func (m *Word) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Word.Merge(m, src)
}
func (m *Word) XXX_Size() int {
	return xxx_messageInfo_Word.Size(m)
}
func (m *Word) XXX_DiscardUnknown() {
	xxx_messageInfo_Word.DiscardUnknown(m)
}

var xxx_messageInfo_Word proto.InternalMessageInfo

func (m *Word) GetWord() string {
	if m != nil {
		return m.Word
	}
	return ""
}

func (m *Word) GetAlphagram() string {
	if m != nil {
		return m.Alphagram
	}
	return ""
}

func (m *Word) GetDefinition() string {
	if m != nil {
		return m.Definition
	}
	return ""
}

func (m *Word) GetFrontHooks() string {
	if m != nil {
		return m.FrontHooks
	}
	return ""
}

func (m *Word) GetBackHooks() string {
	if m != nil {
		return m.BackHooks
	}
	return ""
}

func (m *Word) GetLexiconSymbols() string {
	if m != nil {
		return m.LexiconSymbols
	}
	return ""
}

func (m *Word) GetInnerFrontHook() bool {
	if m != nil {
		return m.InnerFrontHook
	}
	return false
}

func (m *Word) GetInnerBackHook() bool {
	if m != nil {
		return m.InnerBackHook
	}
	return false
}

// A SearchRequest encapsulates a number of varied conditions and lets one
// search for questions.
type SearchRequest struct {
	Searchparams         []*SearchRequest_SearchParam `protobuf:"bytes,1,rep,name=searchparams,proto3" json:"searchparams,omitempty"`
	Expand               bool                         `protobuf:"varint,2,opt,name=expand,proto3" json:"expand,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *SearchRequest) Reset()         { *m = SearchRequest{} }
func (m *SearchRequest) String() string { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()    {}
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_14226a2334c91a98, []int{2}
}

func (m *SearchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchRequest.Unmarshal(m, b)
}
func (m *SearchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchRequest.Marshal(b, m, deterministic)
}
func (m *SearchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchRequest.Merge(m, src)
}
func (m *SearchRequest) XXX_Size() int {
	return xxx_messageInfo_SearchRequest.Size(m)
}
func (m *SearchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchRequest proto.InternalMessageInfo

func (m *SearchRequest) GetSearchparams() []*SearchRequest_SearchParam {
	if m != nil {
		return m.Searchparams
	}
	return nil
}

func (m *SearchRequest) GetExpand() bool {
	if m != nil {
		return m.Expand
	}
	return false
}

type SearchRequest_MinMax struct {
	// Used for length, prob range, prob limit, num anagrams,
	// num_vowels, point value
	Min                  int32    `protobuf:"varint,1,opt,name=min,proto3" json:"min,omitempty"`
	Max                  int32    `protobuf:"varint,2,opt,name=max,proto3" json:"max,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchRequest_MinMax) Reset()         { *m = SearchRequest_MinMax{} }
func (m *SearchRequest_MinMax) String() string { return proto.CompactTextString(m) }
func (*SearchRequest_MinMax) ProtoMessage()    {}
func (*SearchRequest_MinMax) Descriptor() ([]byte, []int) {
	return fileDescriptor_14226a2334c91a98, []int{2, 0}
}

func (m *SearchRequest_MinMax) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchRequest_MinMax.Unmarshal(m, b)
}
func (m *SearchRequest_MinMax) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchRequest_MinMax.Marshal(b, m, deterministic)
}
func (m *SearchRequest_MinMax) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchRequest_MinMax.Merge(m, src)
}
func (m *SearchRequest_MinMax) XXX_Size() int {
	return xxx_messageInfo_SearchRequest_MinMax.Size(m)
}
func (m *SearchRequest_MinMax) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchRequest_MinMax.DiscardUnknown(m)
}

var xxx_messageInfo_SearchRequest_MinMax proto.InternalMessageInfo

func (m *SearchRequest_MinMax) GetMin() int32 {
	if m != nil {
		return m.Min
	}
	return 0
}

func (m *SearchRequest_MinMax) GetMax() int32 {
	if m != nil {
		return m.Max
	}
	return 0
}

type SearchRequest_StringValue struct {
	// Used for lexicon, matching anagram, not_in_lexicon
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchRequest_StringValue) Reset()         { *m = SearchRequest_StringValue{} }
func (m *SearchRequest_StringValue) String() string { return proto.CompactTextString(m) }
func (*SearchRequest_StringValue) ProtoMessage()    {}
func (*SearchRequest_StringValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_14226a2334c91a98, []int{2, 1}
}

func (m *SearchRequest_StringValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchRequest_StringValue.Unmarshal(m, b)
}
func (m *SearchRequest_StringValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchRequest_StringValue.Marshal(b, m, deterministic)
}
func (m *SearchRequest_StringValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchRequest_StringValue.Merge(m, src)
}
func (m *SearchRequest_StringValue) XXX_Size() int {
	return xxx_messageInfo_SearchRequest_StringValue.Size(m)
}
func (m *SearchRequest_StringValue) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchRequest_StringValue.DiscardUnknown(m)
}

var xxx_messageInfo_SearchRequest_StringValue proto.InternalMessageInfo

func (m *SearchRequest_StringValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type SearchRequest_StringArray struct {
	// Used for alphagram_list
	Values               []string `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchRequest_StringArray) Reset()         { *m = SearchRequest_StringArray{} }
func (m *SearchRequest_StringArray) String() string { return proto.CompactTextString(m) }
func (*SearchRequest_StringArray) ProtoMessage()    {}
func (*SearchRequest_StringArray) Descriptor() ([]byte, []int) {
	return fileDescriptor_14226a2334c91a98, []int{2, 2}
}

func (m *SearchRequest_StringArray) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchRequest_StringArray.Unmarshal(m, b)
}
func (m *SearchRequest_StringArray) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchRequest_StringArray.Marshal(b, m, deterministic)
}
func (m *SearchRequest_StringArray) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchRequest_StringArray.Merge(m, src)
}
func (m *SearchRequest_StringArray) XXX_Size() int {
	return xxx_messageInfo_SearchRequest_StringArray.Size(m)
}
func (m *SearchRequest_StringArray) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchRequest_StringArray.DiscardUnknown(m)
}

var xxx_messageInfo_SearchRequest_StringArray proto.InternalMessageInfo

func (m *SearchRequest_StringArray) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

type SearchRequest_NumberArray struct {
	// Used for prob list
	Values               []int32  `protobuf:"varint,1,rep,packed,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchRequest_NumberArray) Reset()         { *m = SearchRequest_NumberArray{} }
func (m *SearchRequest_NumberArray) String() string { return proto.CompactTextString(m) }
func (*SearchRequest_NumberArray) ProtoMessage()    {}
func (*SearchRequest_NumberArray) Descriptor() ([]byte, []int) {
	return fileDescriptor_14226a2334c91a98, []int{2, 3}
}

func (m *SearchRequest_NumberArray) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchRequest_NumberArray.Unmarshal(m, b)
}
func (m *SearchRequest_NumberArray) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchRequest_NumberArray.Marshal(b, m, deterministic)
}
func (m *SearchRequest_NumberArray) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchRequest_NumberArray.Merge(m, src)
}
func (m *SearchRequest_NumberArray) XXX_Size() int {
	return xxx_messageInfo_SearchRequest_NumberArray.Size(m)
}
func (m *SearchRequest_NumberArray) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchRequest_NumberArray.DiscardUnknown(m)
}

var xxx_messageInfo_SearchRequest_NumberArray proto.InternalMessageInfo

func (m *SearchRequest_NumberArray) GetValues() []int32 {
	if m != nil {
		return m.Values
	}
	return nil
}

type SearchRequest_NumberValue struct {
	Value                int32    `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchRequest_NumberValue) Reset()         { *m = SearchRequest_NumberValue{} }
func (m *SearchRequest_NumberValue) String() string { return proto.CompactTextString(m) }
func (*SearchRequest_NumberValue) ProtoMessage()    {}
func (*SearchRequest_NumberValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_14226a2334c91a98, []int{2, 4}
}

func (m *SearchRequest_NumberValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchRequest_NumberValue.Unmarshal(m, b)
}
func (m *SearchRequest_NumberValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchRequest_NumberValue.Marshal(b, m, deterministic)
}
func (m *SearchRequest_NumberValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchRequest_NumberValue.Merge(m, src)
}
func (m *SearchRequest_NumberValue) XXX_Size() int {
	return xxx_messageInfo_SearchRequest_NumberValue.Size(m)
}
func (m *SearchRequest_NumberValue) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchRequest_NumberValue.DiscardUnknown(m)
}

var xxx_messageInfo_SearchRequest_NumberValue proto.InternalMessageInfo

func (m *SearchRequest_NumberValue) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type SearchRequest_SearchParam struct {
	Condition SearchRequest_Condition `protobuf:"varint,1,opt,name=condition,proto3,enum=wordsearcher.SearchRequest_Condition" json:"condition,omitempty"`
	// Types that are valid to be assigned to Conditionparam:
	//	*SearchRequest_SearchParam_Minmax
	//	*SearchRequest_SearchParam_Stringvalue
	//	*SearchRequest_SearchParam_Stringarray
	//	*SearchRequest_SearchParam_Numberarray
	//	*SearchRequest_SearchParam_Numbervalue
	Conditionparam       isSearchRequest_SearchParam_Conditionparam `protobuf_oneof:"conditionparam"`
	XXX_NoUnkeyedLiteral struct{}                                   `json:"-"`
	XXX_unrecognized     []byte                                     `json:"-"`
	XXX_sizecache        int32                                      `json:"-"`
}

func (m *SearchRequest_SearchParam) Reset()         { *m = SearchRequest_SearchParam{} }
func (m *SearchRequest_SearchParam) String() string { return proto.CompactTextString(m) }
func (*SearchRequest_SearchParam) ProtoMessage()    {}
func (*SearchRequest_SearchParam) Descriptor() ([]byte, []int) {
	return fileDescriptor_14226a2334c91a98, []int{2, 5}
}

func (m *SearchRequest_SearchParam) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchRequest_SearchParam.Unmarshal(m, b)
}
func (m *SearchRequest_SearchParam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchRequest_SearchParam.Marshal(b, m, deterministic)
}
func (m *SearchRequest_SearchParam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchRequest_SearchParam.Merge(m, src)
}
func (m *SearchRequest_SearchParam) XXX_Size() int {
	return xxx_messageInfo_SearchRequest_SearchParam.Size(m)
}
func (m *SearchRequest_SearchParam) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchRequest_SearchParam.DiscardUnknown(m)
}

var xxx_messageInfo_SearchRequest_SearchParam proto.InternalMessageInfo

func (m *SearchRequest_SearchParam) GetCondition() SearchRequest_Condition {
	if m != nil {
		return m.Condition
	}
	return SearchRequest_LEXICON
}

type isSearchRequest_SearchParam_Conditionparam interface {
	isSearchRequest_SearchParam_Conditionparam()
}

type SearchRequest_SearchParam_Minmax struct {
	Minmax *SearchRequest_MinMax `protobuf:"bytes,2,opt,name=minmax,proto3,oneof"`
}

type SearchRequest_SearchParam_Stringvalue struct {
	Stringvalue *SearchRequest_StringValue `protobuf:"bytes,3,opt,name=stringvalue,proto3,oneof"`
}

type SearchRequest_SearchParam_Stringarray struct {
	Stringarray *SearchRequest_StringArray `protobuf:"bytes,4,opt,name=stringarray,proto3,oneof"`
}

type SearchRequest_SearchParam_Numberarray struct {
	Numberarray *SearchRequest_NumberArray `protobuf:"bytes,5,opt,name=numberarray,proto3,oneof"`
}

type SearchRequest_SearchParam_Numbervalue struct {
	Numbervalue *SearchRequest_NumberValue `protobuf:"bytes,6,opt,name=numbervalue,proto3,oneof"`
}

func (*SearchRequest_SearchParam_Minmax) isSearchRequest_SearchParam_Conditionparam() {}

func (*SearchRequest_SearchParam_Stringvalue) isSearchRequest_SearchParam_Conditionparam() {}

func (*SearchRequest_SearchParam_Stringarray) isSearchRequest_SearchParam_Conditionparam() {}

func (*SearchRequest_SearchParam_Numberarray) isSearchRequest_SearchParam_Conditionparam() {}

func (*SearchRequest_SearchParam_Numbervalue) isSearchRequest_SearchParam_Conditionparam() {}

func (m *SearchRequest_SearchParam) GetConditionparam() isSearchRequest_SearchParam_Conditionparam {
	if m != nil {
		return m.Conditionparam
	}
	return nil
}

func (m *SearchRequest_SearchParam) GetMinmax() *SearchRequest_MinMax {
	if x, ok := m.GetConditionparam().(*SearchRequest_SearchParam_Minmax); ok {
		return x.Minmax
	}
	return nil
}

func (m *SearchRequest_SearchParam) GetStringvalue() *SearchRequest_StringValue {
	if x, ok := m.GetConditionparam().(*SearchRequest_SearchParam_Stringvalue); ok {
		return x.Stringvalue
	}
	return nil
}

func (m *SearchRequest_SearchParam) GetStringarray() *SearchRequest_StringArray {
	if x, ok := m.GetConditionparam().(*SearchRequest_SearchParam_Stringarray); ok {
		return x.Stringarray
	}
	return nil
}

func (m *SearchRequest_SearchParam) GetNumberarray() *SearchRequest_NumberArray {
	if x, ok := m.GetConditionparam().(*SearchRequest_SearchParam_Numberarray); ok {
		return x.Numberarray
	}
	return nil
}

func (m *SearchRequest_SearchParam) GetNumbervalue() *SearchRequest_NumberValue {
	if x, ok := m.GetConditionparam().(*SearchRequest_SearchParam_Numbervalue); ok {
		return x.Numbervalue
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SearchRequest_SearchParam) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SearchRequest_SearchParam_Minmax)(nil),
		(*SearchRequest_SearchParam_Stringvalue)(nil),
		(*SearchRequest_SearchParam_Stringarray)(nil),
		(*SearchRequest_SearchParam_Numberarray)(nil),
		(*SearchRequest_SearchParam_Numbervalue)(nil),
	}
}

type SearchResponse struct {
	Alphagrams           []*Alphagram `protobuf:"bytes,1,rep,name=alphagrams,proto3" json:"alphagrams,omitempty"`
	Lexicon              string       `protobuf:"bytes,2,opt,name=lexicon,proto3" json:"lexicon,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *SearchResponse) Reset()         { *m = SearchResponse{} }
func (m *SearchResponse) String() string { return proto.CompactTextString(m) }
func (*SearchResponse) ProtoMessage()    {}
func (*SearchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_14226a2334c91a98, []int{3}
}

func (m *SearchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchResponse.Unmarshal(m, b)
}
func (m *SearchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchResponse.Marshal(b, m, deterministic)
}
func (m *SearchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchResponse.Merge(m, src)
}
func (m *SearchResponse) XXX_Size() int {
	return xxx_messageInfo_SearchResponse.Size(m)
}
func (m *SearchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SearchResponse proto.InternalMessageInfo

func (m *SearchResponse) GetAlphagrams() []*Alphagram {
	if m != nil {
		return m.Alphagrams
	}
	return nil
}

func (m *SearchResponse) GetLexicon() string {
	if m != nil {
		return m.Lexicon
	}
	return ""
}

func init() {
	proto.RegisterEnum("wordsearcher.SearchRequest_Condition", SearchRequest_Condition_name, SearchRequest_Condition_value)
	proto.RegisterEnum("wordsearcher.SearchRequest_NotInLexCondition", SearchRequest_NotInLexCondition_name, SearchRequest_NotInLexCondition_value)
	proto.RegisterType((*Alphagram)(nil), "wordsearcher.Alphagram")
	proto.RegisterType((*Word)(nil), "wordsearcher.Word")
	proto.RegisterType((*SearchRequest)(nil), "wordsearcher.SearchRequest")
	proto.RegisterType((*SearchRequest_MinMax)(nil), "wordsearcher.SearchRequest.MinMax")
	proto.RegisterType((*SearchRequest_StringValue)(nil), "wordsearcher.SearchRequest.StringValue")
	proto.RegisterType((*SearchRequest_StringArray)(nil), "wordsearcher.SearchRequest.StringArray")
	proto.RegisterType((*SearchRequest_NumberArray)(nil), "wordsearcher.SearchRequest.NumberArray")
	proto.RegisterType((*SearchRequest_NumberValue)(nil), "wordsearcher.SearchRequest.NumberValue")
	proto.RegisterType((*SearchRequest_SearchParam)(nil), "wordsearcher.SearchRequest.SearchParam")
	proto.RegisterType((*SearchResponse)(nil), "wordsearcher.SearchResponse")
}

func init() { proto.RegisterFile("rpc/wordsearcher/searcher.proto", fileDescriptor_14226a2334c91a98) }

var fileDescriptor_14226a2334c91a98 = []byte{
	// 866 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0x4d, 0x6f, 0xdb, 0x46,
	0x10, 0x35, 0x2d, 0x93, 0x16, 0x87, 0x96, 0xbc, 0x1e, 0xa4, 0x29, 0xe1, 0xa6, 0x8d, 0xa0, 0x20,
	0x8d, 0x0e, 0x85, 0x03, 0xb8, 0x87, 0x5e, 0x72, 0xa1, 0x14, 0x46, 0x22, 0x42, 0x91, 0xee, 0x52,
	0x96, 0xd3, 0x13, 0x41, 0x49, 0x8c, 0x4d, 0x44, 0x5a, 0xaa, 0x94, 0xdc, 0xca, 0xbf, 0xa3, 0xe7,
	0xfe, 0xaa, 0xa2, 0xfd, 0x3b, 0x2d, 0x76, 0x97, 0x14, 0xa9, 0xc0, 0x50, 0x73, 0xdb, 0x79, 0xfb,
	0xe6, 0xcd, 0xc7, 0x23, 0x96, 0xf0, 0x3c, 0x5b, 0x4e, 0x5f, 0xff, 0x9e, 0x66, 0xb3, 0x55, 0x1c,
	0x65, 0xd3, 0xbb, 0x38, 0x7b, 0x5d, 0x1c, 0x2e, 0x96, 0x59, 0xba, 0x4e, 0xf1, 0xa4, 0x7a, 0xd9,
	0xfe, 0x5b, 0x01, 0xdd, 0x9a, 0x2f, 0xef, 0xa2, 0xdb, 0x2c, 0x5a, 0xe0, 0x33, 0xd0, 0xa3, 0x22,
	0x30, 0x95, 0x96, 0xd2, 0xd1, 0x69, 0x09, 0x60, 0x07, 0x54, 0x91, 0x6b, 0x1e, 0xb6, 0x6a, 0x1d,
	0xe3, 0x12, 0x2f, 0xaa, 0x4a, 0x17, 0x37, 0x69, 0x36, 0xa3, 0x92, 0x80, 0x6d, 0x38, 0x89, 0x37,
	0xcb, 0x88, 0xcd, 0xe2, 0x19, 0x8d, 0x97, 0x99, 0x59, 0x6b, 0x29, 0x9d, 0x3a, 0xdd, 0xc1, 0xf0,
	0x29, 0x68, 0xf3, 0x98, 0xdd, 0xae, 0xef, 0xcc, 0xa3, 0x96, 0xd2, 0x51, 0x69, 0x1e, 0x61, 0x0b,
	0x8c, 0x65, 0x96, 0x4e, 0xa2, 0x49, 0x32, 0x4f, 0xd6, 0x0f, 0xa6, 0x2a, 0x2e, 0xab, 0x10, 0x57,
	0x9f, 0xa6, 0x8b, 0x49, 0xc2, 0xa2, 0x75, 0x92, 0xb2, 0x95, 0xa9, 0xb5, 0x94, 0x4e, 0x8d, 0xee,
	0x60, 0xed, 0x3f, 0x0e, 0xe1, 0x88, 0x77, 0x84, 0x08, 0x47, 0xbc, 0xa7, 0x7c, 0x1a, 0x71, 0xde,
	0x1d, 0xf3, 0xf0, 0xf3, 0x31, 0xbf, 0x03, 0x98, 0xc5, 0x1f, 0x13, 0x96, 0x70, 0x25, 0xd1, 0xba,
	0x4e, 0x2b, 0x08, 0x3e, 0x07, 0xe3, 0x63, 0x96, 0xb2, 0x75, 0x78, 0x97, 0xa6, 0x9f, 0x56, 0xa2,
	0x7b, 0x9d, 0x82, 0x80, 0x06, 0x1c, 0xc1, 0x6f, 0x01, 0x26, 0xd1, 0xf4, 0x53, 0x7e, 0xaf, 0x4a,
	0x7d, 0x8e, 0xc8, 0xeb, 0x57, 0x70, 0x3a, 0x8f, 0x37, 0xc9, 0x34, 0x65, 0xe1, 0xea, 0x61, 0x31,
	0x49, 0xe7, 0x72, 0x02, 0x9d, 0x36, 0x73, 0x38, 0x90, 0x28, 0x76, 0x80, 0x24, 0x8c, 0xc5, 0x59,
	0x58, 0x96, 0x33, 0x8f, 0xc5, 0x26, 0x9b, 0x02, 0x7f, 0x57, 0x94, 0xc4, 0xef, 0xe1, 0x54, 0x32,
	0xb7, 0x75, 0xcd, 0xba, 0x20, 0x36, 0x04, 0xdc, 0xcd, 0x6b, 0xb7, 0xff, 0xaa, 0x43, 0x23, 0x10,
	0x86, 0xd1, 0xf8, 0xd7, 0xfb, 0x78, 0xb5, 0xc6, 0xf7, 0x70, 0x22, 0x1d, 0x5c, 0x46, 0x59, 0xb4,
	0x58, 0x99, 0x8a, 0xb0, 0xf6, 0xd5, 0xae, 0xb5, 0x3b, 0x29, 0x79, 0x74, 0xc5, 0xf9, 0x74, 0x27,
	0x99, 0x5b, 0x2a, 0x2d, 0x16, 0x4b, 0xad, 0xd3, 0x3c, 0x3a, 0xff, 0x01, 0xb4, 0x61, 0xc2, 0x86,
	0xd1, 0x06, 0x09, 0xd4, 0x16, 0x09, 0x13, 0x66, 0xa8, 0x94, 0x1f, 0x05, 0x12, 0x6d, 0x44, 0x02,
	0x47, 0xa2, 0xcd, 0xf9, 0x0b, 0x30, 0x82, 0x75, 0x96, 0xb0, 0xdb, 0x71, 0x34, 0xbf, 0x8f, 0xf1,
	0x09, 0xa8, 0xbf, 0xf1, 0x43, 0xee, 0xa0, 0x0c, 0xce, 0x5f, 0x16, 0x24, 0x2b, 0xcb, 0xa2, 0x07,
	0x5e, 0x59, 0xe0, 0x72, 0x00, 0x9d, 0xe6, 0x11, 0xa7, 0x79, 0xf7, 0x8b, 0x49, 0x9c, 0x3d, 0x46,
	0x53, 0xb7, 0xb4, 0x17, 0x05, 0xed, 0x91, 0x92, 0x6a, 0x51, 0xf2, 0x9f, 0x1a, 0x18, 0x95, 0xd9,
	0xb1, 0x07, 0xfa, 0x34, 0x65, 0x33, 0xf9, 0x99, 0x70, 0x66, 0xf3, 0xf2, 0xe5, 0xbe, 0xbd, 0xf5,
	0x0a, 0x32, 0x2d, 0xf3, 0xf0, 0x0d, 0x68, 0x8b, 0x84, 0x15, 0x1b, 0x30, 0x2e, 0xdb, 0xfb, 0x14,
	0xe4, 0x12, 0x07, 0x07, 0x34, 0xcf, 0xc1, 0xf7, 0x60, 0xac, 0xc4, 0x16, 0x64, 0xbb, 0x35, 0x21,
	0xb1, 0xdf, 0xbc, 0x72, 0xb3, 0x83, 0x03, 0x5a, 0xcd, 0x2e, 0xc5, 0x22, 0xbe, 0x2b, 0xf1, 0x5d,
	0x7f, 0x91, 0x98, 0x58, 0x6d, 0x29, 0x26, 0xb2, 0xb9, 0x18, 0x13, 0x1b, 0x95, 0x62, 0xea, 0xff,
	0x8b, 0x55, 0x7c, 0xe2, 0x62, 0x95, 0xec, 0x52, 0x4c, 0x8e, 0xa9, 0x7d, 0xa9, 0xd8, 0x76, 0xcc,
	0x4a, 0x76, 0x97, 0x40, 0x73, 0xbb, 0x7e, 0xf1, 0xdd, 0xb6, 0xff, 0x55, 0x40, 0xdf, 0x9a, 0x83,
	0x06, 0x1c, 0xbb, 0xf6, 0x07, 0xa7, 0xe7, 0x7b, 0xe4, 0x00, 0x01, 0x34, 0xd7, 0xf6, 0xfa, 0xa3,
	0x01, 0x51, 0xf0, 0x2b, 0x38, 0xbb, 0xa2, 0x7e, 0xd7, 0xea, 0x3a, 0xae, 0x33, 0xfa, 0x25, 0xa4,
	0x96, 0xd7, 0xb7, 0xc9, 0x21, 0x3e, 0x01, 0x52, 0x85, 0x5d, 0x27, 0x18, 0x91, 0xda, 0xe7, 0x64,
	0xd7, 0x19, 0x3a, 0x23, 0x72, 0x84, 0x4f, 0x01, 0xbd, 0xeb, 0x61, 0xd7, 0xa6, 0xa1, 0xff, 0x2e,
	0xb4, 0x3c, 0xab, 0x4f, 0xad, 0x61, 0x40, 0x54, 0x2e, 0x52, 0xe2, 0x63, 0xff, 0xc6, 0x76, 0x03,
	0xa2, 0xe1, 0x29, 0x18, 0x57, 0xbe, 0xe3, 0x8d, 0xc2, 0xb1, 0xe5, 0x5e, 0xdb, 0xa4, 0xce, 0x69,
	0x43, 0x6b, 0xd4, 0x1b, 0x38, 0x5e, 0xbf, 0xc8, 0x26, 0x3a, 0x22, 0x34, 0x2d, 0xf7, 0x6a, 0x20,
	0x42, 0x59, 0x1f, 0x38, 0xe6, 0xf9, 0xa3, 0xd0, 0xf1, 0xc2, 0x62, 0x18, 0x03, 0x75, 0x50, 0xa9,
	0xdd, 0xb7, 0x3f, 0x90, 0x13, 0x6c, 0x80, 0x7e, 0xe3, 0xd3, 0xb7, 0x92, 0xdd, 0x68, 0xbf, 0x81,
	0x33, 0x2f, 0x5d, 0x3b, 0xcc, 0x8d, 0x37, 0xe5, 0x22, 0xce, 0xa0, 0xe1, 0x8f, 0x06, 0x36, 0x0d,
	0x6d, 0xaf, 0xef, 0x3a, 0xc1, 0x80, 0x1c, 0xc8, 0x59, 0xed, 0xb1, 0xe3, 0x5f, 0x07, 0xe1, 0xd8,
	0xa6, 0x81, 0xe3, 0x7b, 0x44, 0x69, 0x4f, 0xa1, 0x59, 0x6c, 0x7f, 0xb5, 0x4c, 0xd9, 0x2a, 0xc6,
	0x9f, 0x00, 0xb6, 0xef, 0x69, 0xf1, 0xa6, 0x7c, 0xbd, 0xeb, 0xd7, 0xf6, 0xa7, 0x43, 0x2b, 0x54,
	0x34, 0xe1, 0x38, 0x7f, 0x04, 0xf3, 0x77, 0xb9, 0x08, 0x2f, 0xff, 0x54, 0x80, 0xfc, 0xcc, 0xbd,
	0x4d, 0x52, 0x16, 0xe4, 0x22, 0xd8, 0x03, 0x4d, 0x9e, 0xf1, 0x9b, 0x3d, 0x5f, 0xc3, 0xf9, 0xb3,
	0xc7, 0x2f, 0xf3, 0x66, 0xdf, 0x82, 0x66, 0x8b, 0x77, 0x0a, 0xf7, 0xf2, 0xf6, 0xab, 0x4c, 0x34,
	0xf1, 0x77, 0xfd, 0xf1, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x86, 0xc4, 0x54, 0xd1, 0x80, 0x07,
	0x00, 0x00,
}
