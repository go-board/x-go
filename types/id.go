package types

import (
	"encoding/json"
	"strconv"
	"strings"

	"golang.org/x/xerrors"
)

var idEncoder func(ID) string = func(i ID) string { return strconv.FormatInt(int64(i), 10) }
var idDecoder func(string) (ID, error) = func(s string) (ID, error) {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return ID(i), err
}

func Init(encode func(ID) string, decode func(string) (ID, error)) {
	idEncoder = encode
	idDecoder = decode
}

// ID is int64, useful for some convenient transform case.
// Example:
//	json api, to protect real id, developer can embed this
//	ID into a struct to do automatically transform int64 to
//	string with encrypt.
type ID int64

func (i ID) MarshalText() (text []byte, err error) {
	if i < 0 {
		return nil, xerrors.Errorf("err: ID(%d) must greater than 0", i)
	}
	str := idEncoder(i)
	return []byte(str), nil
}

func (i *ID) UnmarshalText(text []byte) error {
	id, err := idDecoder(string(text))
	if err != nil {
		return xerrors.Errorf("unmarshal id failed, %w", err)
	}

	*i = ID(id)
	return nil
}

func (i ID) MarshalJSON() ([]byte, error) {
	text, err := i.MarshalText()
	if err != nil {
		return nil, err
	}
	return json.Marshal(string(text))
}

func (i *ID) UnmarshalJSON(data []byte) error {
	var text string
	err := json.Unmarshal(data, &text)
	if err != nil {
		return err
	}
	return i.UnmarshalText([]byte(text))
}

func ParseIDList(str string) ([]ID, error) {
	items := strings.Split(str, ",")
	ids := make([]ID, 0, len(items))
	for _, item := range items {
		id := new(ID)
		err := id.UnmarshalText([]byte(item))
		if err != nil {
			continue
		}
		ids = append(ids, *id)
	}
	return ids, nil
}

func ParseInt64List(str string) ([]int64, error) {
	ids, err := ParseIDList(str)
	if err != nil {
		return nil, err
	}
	int64s := make([]int64, 0, len(ids))
	for _, id := range ids {
		int64s = append(int64s, int64(id))
	}
	return int64s, nil
}

func FormatIDList(ids []ID) (string, error) {
	items := make([]string, 0, len(ids))
	for _, id := range ids {
		item, err := id.MarshalJSON()
		if err != nil {
			continue
		}
		items = append(items, string(item))
	}
	return strings.Join(items, ","), nil
}
