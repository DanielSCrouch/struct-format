package format_test

import (
	"format_test/pkg/utils/format"
	"reflect"
	"testing"
)

// TestFormatted
func TestFormatted(t *testing.T) {

	item1 := &Cell{Identity: &Identifier{Guid: "foo", Alias: "baaar"}, CellStatus: CellStatus_AVAILABLE, NodePort: 30002, ClusterIdentifier: &Identifier{Guid: "foo-Cluster", Alias: "baaar-Cluster"}}
	item2 := &Cell{Identity: &Identifier{Guid: "fooooo", Alias: "bar"}, CellStatus: CellStatus_BROKEN, NodePort: 30005, ClusterIdentifier: &Identifier{Guid: "foooooo-Cluster", Alias: "baaaaar-Cluster"}}
	item3 := &Cell{Identity: &Identifier{Guid: "fooo", Alias: "bar"}, CellStatus: CellStatus_NOT_CONNECTED, NodePort: 30004, ClusterIdentifier: &Identifier{Guid: "fo-Cluster", Alias: "bar-Cluster"}}
	itemList := []*Cell{item1, item2, item3}
	fieldNames := []string{"ID", "Name", "Status", "Node Port", "Cluster"}
	fieldPaths := []string{"Identity.Guid", "Identity.Alias", "CellStatus", "NodePort", "ClusterIdentifier.Alias"}

	_, err := format.FormattedList(itemList, fieldNames, fieldPaths)
	if err != nil {
		t.Fatal(err)
	}

}

// TestGetFieldValue
func TestGetFieldValue(t *testing.T) {

	testcases := []struct {
		item        interface{}
		fieldPath   []string
		expected    interface{}
		expectedErr error
	}{
		{ // Simple struct
			item: struct {
				NestedValue struct{ SecondNestedValue string }
			}{
				NestedValue: struct{ SecondNestedValue string }{
					"foo",
				},
			},
			fieldPath:   []string{"NestedValue", "SecondNestedValue"},
			expected:    "foo",
			expectedErr: nil,
		},
		{ // Struct with pointers
			item: struct {
				NestedValue *struct{ SecondNestedValue string }
			}{
				NestedValue: &struct{ SecondNestedValue string }{
					"foo",
				},
			},
			fieldPath:   []string{"NestedValue", "SecondNestedValue"},
			expected:    "foo",
			expectedErr: nil,
		},
		{ // Invalid field path
			item: struct {
				NestedValue *struct{ SecondNestedValue string }
			}{
				NestedValue: &struct{ SecondNestedValue string }{
					"foo",
				},
			},
			fieldPath:   []string{"NestedValue", "X"},
			expected:    nil,
			expectedErr: &format.InvalidFieldPath{},
		},
		{ // Struct contains invalid Slice
			item: struct {
				NestedValue []*struct{ SecondNestedValue string }
			}{
				NestedValue: []*struct{ SecondNestedValue string }{
					{
						SecondNestedValue: "foo",
					},
				},
			},
			fieldPath:   []string{"NestedValue", "SecondNestedValue"},
			expected:    nil,
			expectedErr: &format.InvalidFieldPathSlice{},
		},
	}

	for i, testcase := range testcases {
		value, err := format.GetFieldValue(testcase.item, testcase.fieldPath)
		if err != nil && reflect.TypeOf(err) != reflect.TypeOf(testcase.expectedErr) {
			t.Fatal(err)
		}

		if value != testcase.expected {
			t.Fatalf("failed test case %d", i+1)
		}

	}
}
