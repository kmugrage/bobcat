package dsl

import (
	. "github.com/ThoughtWorksStudios/bobcat/test_helpers"
	"log"
	"testing"
	"time"
)

func TestSearchNodesWhenGivenSliceOfNodes(t *testing.T) {
	node1 := Node{Kind: "integer", Name: "one", Value: 1}
	node2 := Node{Kind: "integer", Name: "two", Value: 2}
	expected := NodeSet{node1, node2}
	actual := searchNodes([]interface{}{node1, node2})
	AssertEqual(t, expected.String(), actual.String())
}

func TestSearchNodesReturnsEmptyNodeSetWhenReceivesNil(t *testing.T) {
	expected := NodeSet{}
	actual := searchNodes(nil)
	AssertEqual(t, expected.String(), actual.String())
}

func TestSearchNodesWhenGivenListOfNonNodes(t *testing.T) {
	node1 := Node{Kind: "string", Name: "thing", Value: "blah"}
	node2 := Node{Kind: "integer", Name: "value", Value: 42}
	node3 := Node{Kind: "dict", Name: "city", Value: "city"}
	expected := NodeSet{node1, node2, node3}
	weirdArgs := []interface{}{[]interface{}{node1, node2, node3}}
	actual := searchNodes(weirdArgs)
	AssertEqual(t, expected.String(), actual.String())
}

func TestSearchNodesWhenGivenListOfNodesAndValues(t *testing.T) {
	topNode := Node{Kind: "string", Name: "thing", Value: "blah"}
	node1 := Node{Kind: "integer", Name: "value", Value: 42}
	node2 := Node{Kind: "dict", Name: "city", Value: "city"}
	expected := NodeSet{topNode, node1, node2}
	weirdArgs := []interface{}{topNode, []interface{}{node1, node2}}
	actual := searchNodes(weirdArgs)
	AssertEqual(t, expected.String(), actual.String())
}

func TestDelimitedNodeSliceWhereFirstAndRestAreNodes(t *testing.T) {
	first := Node{Kind: "string", Name: "thing", Value: "blah"}
	n := Node{Kind: "integer", Name: "value", Value: 42}
	var rest interface{} = []interface{}{n}
	expected := NodeSet{first, n}
	actual := delimitedNodeSlice(first, rest)
	AssertEqual(t, expected.String(), actual.String())
}

func TestDelimitedNodeSliceWhereRestIsSliceOfNodes(t *testing.T) {
	first := Node{Kind: "string", Name: "thing", Value: "blah"}
	node1 := Node{Kind: "integer", Name: "value", Value: 42}
	node2 := Node{Kind: "dict", Name: "city", Value: "city"}
	expected := NodeSet{first, node1, node2}
	rest := []interface{}{[]interface{}{node1, node2}}
	actual := delimitedNodeSlice(first, rest)
	AssertEqual(t, expected.String(), actual.String())
}
func TestDelimitedNodeSliceWhereRestIsComplex(t *testing.T) {
	first := Node{Kind: "string", Name: "thing", Value: "blah"}
	node2 := Node{Kind: "integer", Name: "value", Value: 42}
	node3 := Node{Kind: "dict", Name: "city", Value: "city"}
	node4 := Node{Kind: "decimal", Name: "age"}
	expected := NodeSet{first, node2, node3, node4}
	rest := []interface{}{node2, []interface{}{node3, node4}}
	actual := delimitedNodeSlice(first, rest)
	AssertEqual(t, expected.String(), actual.String())
}

func TestParseDateLikeJS(t *testing.T) {
	specs := map[string]time.Time{
		"2017-07-11":                parse("2017-07-11 00:00:00 +0000"),
		"2017-07-11T00:14:56":       parse("2017-07-11 00:14:56 +0000"),
		"2017-07-11T00:14:56Z":      parse("2017-07-11 00:14:56 +0000"),
		"2017-07-11T00:14:56-0730":  parse("2017-07-11 00:14:56 -0730"),
		"2017-07-11T00:14:56-08:30": parse("2017-07-11 00:14:56 -0830"),
	}

	for ts, expected := range specs {
		actual, err := ParseDateLikeJS(ts)
		AssertNil(t, err, "Got an error while parsing date: %v", err)
		AssertTimeEqual(t, expected, actual)
	}
}

func TestParseDateLikeJSReturnsError(t *testing.T) {
	input := "2017-07-19T13:00:00Z-700"
	expected := "Not a parsable timestamp: 2017-07-19T13:00:00Z-700"
	_, err := ParseDateLikeJS(input)
	ExpectsError(t, expected, err)
}

func TestDefaultToEmptySlice(t *testing.T) {
	expected := NodeSet{}
	actual := defaultToEmptySlice(nil)
	AssertEqual(t, expected.String(), actual.String())

	node1 := Node{Kind: "integer", Name: "one", Value: 1}
	node2 := Node{Kind: "integer", Name: "two", Value: 2}
	expected = NodeSet{node1, node2}
	actual = defaultToEmptySlice(expected)
	AssertEqual(t, expected.String(), actual.String())
}

func parse(stamp string) time.Time {
	t, e := time.Parse("2006-01-02 15:04:05 -0700", stamp)
	if e != nil {
		log.Fatalf("error? %v", e)
	}
	return t
}
