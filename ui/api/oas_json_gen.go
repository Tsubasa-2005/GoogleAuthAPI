// Code generated by ogen, DO NOT EDIT.

package api

import (
	"math/bits"
	"strconv"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"

	"github.com/ogen-go/ogen/validate"
)

// Encode encodes Email as json.
func (s Email) Encode(e *jx.Encoder) {
	unwrapped := string(s)

	e.Str(unwrapped)
}

// Decode decodes Email from json.
func (s *Email) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Email to nil")
	}
	var unwrapped string
	if err := func() error {
		v, err := d.Str()
		unwrapped = string(v)
		if err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = Email(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s Email) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Email) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *GetUsersOK) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *GetUsersOK) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("totalPages")
		e.Int(s.TotalPages)
	}
}

var jsonFieldsNameOfGetUsersOK = [1]string{
	0: "totalPages",
}

// Decode decodes GetUsersOK from json.
func (s *GetUsersOK) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode GetUsersOK to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "totalPages":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Int()
				s.TotalPages = int(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"totalPages\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode GetUsersOK")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000001,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfGetUsersOK) {
					name = jsonFieldsNameOfGetUsersOK[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *GetUsersOK) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *GetUsersOK) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes ImageURL as json.
func (s ImageURL) Encode(e *jx.Encoder) {
	unwrapped := string(s)

	e.Str(unwrapped)
}

// Decode decodes ImageURL from json.
func (s *ImageURL) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ImageURL to nil")
	}
	var unwrapped string
	if err := func() error {
		v, err := d.Str()
		unwrapped = string(v)
		if err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = ImageURL(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s ImageURL) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ImageURL) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes Name as json.
func (s Name) Encode(e *jx.Encoder) {
	unwrapped := string(s)

	e.Str(unwrapped)
}

// Decode decodes Name from json.
func (s *Name) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Name to nil")
	}
	var unwrapped string
	if err := func() error {
		v, err := d.Str()
		unwrapped = string(v)
		if err != nil {
			return err
		}
		return nil
	}(); err != nil {
		return errors.Wrap(err, "alias")
	}
	*s = Name(unwrapped)
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s Name) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Name) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *SignInRequest) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *SignInRequest) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("email")
		s.Email.Encode(e)
	}
	{
		e.FieldStart("name")
		s.Name.Encode(e)
	}
	{
		e.FieldStart("imageURL")
		s.ImageURL.Encode(e)
	}
	{
		e.FieldStart("token")
		e.Str(s.Token)
	}
}

var jsonFieldsNameOfSignInRequest = [4]string{
	0: "email",
	1: "name",
	2: "imageURL",
	3: "token",
}

// Decode decodes SignInRequest from json.
func (s *SignInRequest) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode SignInRequest to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "email":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				if err := s.Email.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"email\"")
			}
		case "name":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				if err := s.Name.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"name\"")
			}
		case "imageURL":
			requiredBitSet[0] |= 1 << 2
			if err := func() error {
				if err := s.ImageURL.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"imageURL\"")
			}
		case "token":
			requiredBitSet[0] |= 1 << 3
			if err := func() error {
				v, err := d.Str()
				s.Token = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"token\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode SignInRequest")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00001111,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfSignInRequest) {
					name = jsonFieldsNameOfSignInRequest[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *SignInRequest) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *SignInRequest) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *Unauthorized) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *Unauthorized) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("message")
		e.Str(s.Message)
	}
}

var jsonFieldsNameOfUnauthorized = [1]string{
	0: "message",
}

// Decode decodes Unauthorized from json.
func (s *Unauthorized) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode Unauthorized to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "message":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.Message = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"message\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode Unauthorized")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000001,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfUnauthorized) {
					name = jsonFieldsNameOfUnauthorized[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *Unauthorized) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *Unauthorized) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}