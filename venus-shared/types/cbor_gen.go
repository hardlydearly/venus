// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

package types

import (
	"fmt"
	"io"
	"math"
	"sort"

	abi "github.com/filecoin-project/go-state-types/abi"
	paych "github.com/filecoin-project/go-state-types/builtin/v8/paych"
	crypto "github.com/filecoin-project/go-state-types/crypto"
	proof "github.com/filecoin-project/go-state-types/proof"
	cid "github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

var _ = xerrors.Errorf
var _ = cid.Undef
var _ = math.E
var _ = sort.Sort

var lengthBufBlockHeader = []byte{144}

func (t *BlockHeader) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufBlockHeader); err != nil {
		return err
	}

	// t.Miner (address.Address) (struct)
	if err := t.Miner.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.Ticket (types.Ticket) (struct)
	if err := t.Ticket.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.ElectionProof (types.ElectionProof) (struct)
	if err := t.ElectionProof.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.BeaconEntries ([]types.BeaconEntry) (slice)
	if len(t.BeaconEntries) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.BeaconEntries was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajArray, uint64(len(t.BeaconEntries))); err != nil {
		return err
	}
	for _, v := range t.BeaconEntries {
		if err := v.MarshalCBOR(cw); err != nil {
			return err
		}
	}

	// t.WinPoStProof ([]proof.PoStProof) (slice)
	if len(t.WinPoStProof) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.WinPoStProof was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajArray, uint64(len(t.WinPoStProof))); err != nil {
		return err
	}
	for _, v := range t.WinPoStProof {
		if err := v.MarshalCBOR(cw); err != nil {
			return err
		}
	}

	// t.Parents ([]cid.Cid) (slice)
	if len(t.Parents) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.Parents was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajArray, uint64(len(t.Parents))); err != nil {
		return err
	}
	for _, v := range t.Parents {
		if err := cbg.WriteCid(w, v); err != nil {
			return xerrors.Errorf("failed writing cid field t.Parents: %w", err)
		}
	}

	// t.ParentWeight (big.Int) (struct)
	if err := t.ParentWeight.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.Height (abi.ChainEpoch) (int64)
	if t.Height >= 0 {
		if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.Height)); err != nil {
			return err
		}
	} else {
		if err := cw.WriteMajorTypeHeader(cbg.MajNegativeInt, uint64(-t.Height-1)); err != nil {
			return err
		}
	}

	// t.ParentStateRoot (cid.Cid) (struct)

	if err := cbg.WriteCid(cw, t.ParentStateRoot); err != nil {
		return xerrors.Errorf("failed to write cid field t.ParentStateRoot: %w", err)
	}

	// t.ParentMessageReceipts (cid.Cid) (struct)

	if err := cbg.WriteCid(cw, t.ParentMessageReceipts); err != nil {
		return xerrors.Errorf("failed to write cid field t.ParentMessageReceipts: %w", err)
	}

	// t.Messages (cid.Cid) (struct)

	if err := cbg.WriteCid(cw, t.Messages); err != nil {
		return xerrors.Errorf("failed to write cid field t.Messages: %w", err)
	}

	// t.BLSAggregate (crypto.Signature) (struct)
	if err := t.BLSAggregate.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.Timestamp (uint64) (uint64)

	if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.Timestamp)); err != nil {
		return err
	}

	// t.BlockSig (crypto.Signature) (struct)
	if err := t.BlockSig.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.ForkSignaling (uint64) (uint64)

	if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.ForkSignaling)); err != nil {
		return err
	}

	// t.ParentBaseFee (big.Int) (struct)
	if err := t.ParentBaseFee.MarshalCBOR(cw); err != nil {
		return err
	}
	return nil
}

func (t *BlockHeader) UnmarshalCBOR(r io.Reader) (err error) {
	*t = BlockHeader{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 16 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Miner (address.Address) (struct)

	{

		if err := t.Miner.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.Miner: %w", err)
		}

	}
	// t.Ticket (types.Ticket) (struct)

	{

		b, err := cr.ReadByte()
		if err != nil {
			return err
		}
		if b != cbg.CborNull[0] {
			if err := cr.UnreadByte(); err != nil {
				return err
			}
			t.Ticket = new(Ticket)
			if err := t.Ticket.UnmarshalCBOR(cr); err != nil {
				return xerrors.Errorf("unmarshaling t.Ticket pointer: %w", err)
			}
		}

	}
	// t.ElectionProof (types.ElectionProof) (struct)

	{

		b, err := cr.ReadByte()
		if err != nil {
			return err
		}
		if b != cbg.CborNull[0] {
			if err := cr.UnreadByte(); err != nil {
				return err
			}
			t.ElectionProof = new(ElectionProof)
			if err := t.ElectionProof.UnmarshalCBOR(cr); err != nil {
				return xerrors.Errorf("unmarshaling t.ElectionProof pointer: %w", err)
			}
		}

	}
	// t.BeaconEntries ([]types.BeaconEntry) (slice)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.BeaconEntries: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.BeaconEntries = make([]BeaconEntry, extra)
	}

	for i := 0; i < int(extra); i++ {

		var v BeaconEntry
		if err := v.UnmarshalCBOR(cr); err != nil {
			return err
		}

		t.BeaconEntries[i] = v
	}

	// t.WinPoStProof ([]proof.PoStProof) (slice)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.WinPoStProof: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.WinPoStProof = make([]proof.PoStProof, extra)
	}

	for i := 0; i < int(extra); i++ {

		var v proof.PoStProof
		if err := v.UnmarshalCBOR(cr); err != nil {
			return err
		}

		t.WinPoStProof[i] = v
	}

	// t.Parents ([]cid.Cid) (slice)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.Parents: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.Parents = make([]cid.Cid, extra)
	}

	for i := 0; i < int(extra); i++ {

		c, err := cbg.ReadCid(cr)
		if err != nil {
			return xerrors.Errorf("reading cid field t.Parents failed: %w", err)
		}
		t.Parents[i] = c
	}

	// t.ParentWeight (big.Int) (struct)

	{

		if err := t.ParentWeight.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.ParentWeight: %w", err)
		}

	}
	// t.Height (abi.ChainEpoch) (int64)
	{
		maj, extra, err := cr.ReadHeader()
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative oveflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.Height = abi.ChainEpoch(extraI)
	}
	// t.ParentStateRoot (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(cr)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.ParentStateRoot: %w", err)
		}

		t.ParentStateRoot = c

	}
	// t.ParentMessageReceipts (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(cr)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.ParentMessageReceipts: %w", err)
		}

		t.ParentMessageReceipts = c

	}
	// t.Messages (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(cr)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.Messages: %w", err)
		}

		t.Messages = c

	}
	// t.BLSAggregate (crypto.Signature) (struct)

	{

		b, err := cr.ReadByte()
		if err != nil {
			return err
		}
		if b != cbg.CborNull[0] {
			if err := cr.UnreadByte(); err != nil {
				return err
			}
			t.BLSAggregate = new(crypto.Signature)
			if err := t.BLSAggregate.UnmarshalCBOR(cr); err != nil {
				return xerrors.Errorf("unmarshaling t.BLSAggregate pointer: %w", err)
			}
		}

	}
	// t.Timestamp (uint64) (uint64)

	{

		maj, extra, err = cr.ReadHeader()
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.Timestamp = uint64(extra)

	}
	// t.BlockSig (crypto.Signature) (struct)

	{

		b, err := cr.ReadByte()
		if err != nil {
			return err
		}
		if b != cbg.CborNull[0] {
			if err := cr.UnreadByte(); err != nil {
				return err
			}
			t.BlockSig = new(crypto.Signature)
			if err := t.BlockSig.UnmarshalCBOR(cr); err != nil {
				return xerrors.Errorf("unmarshaling t.BlockSig pointer: %w", err)
			}
		}

	}
	// t.ForkSignaling (uint64) (uint64)

	{

		maj, extra, err = cr.ReadHeader()
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.ForkSignaling = uint64(extra)

	}
	// t.ParentBaseFee (big.Int) (struct)

	{

		if err := t.ParentBaseFee.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.ParentBaseFee: %w", err)
		}

	}
	return nil
}

var lengthBufTicket = []byte{129}

func (t *Ticket) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufTicket); err != nil {
		return err
	}

	// t.VRFProof (types.VRFPi) (slice)
	if len(t.VRFProof) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.VRFProof was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajByteString, uint64(len(t.VRFProof))); err != nil {
		return err
	}

	if _, err := cw.Write(t.VRFProof[:]); err != nil {
		return err
	}
	return nil
}

func (t *Ticket) UnmarshalCBOR(r io.Reader) (err error) {
	*t = Ticket{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 1 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.VRFProof (types.VRFPi) (slice)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.VRFProof: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.VRFProof = make([]uint8, extra)
	}

	if _, err := io.ReadFull(cr, t.VRFProof[:]); err != nil {
		return err
	}
	return nil
}

var lengthBufElectionProof = []byte{130}

func (t *ElectionProof) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufElectionProof); err != nil {
		return err
	}

	// t.WinCount (int64) (int64)
	if t.WinCount >= 0 {
		if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.WinCount)); err != nil {
			return err
		}
	} else {
		if err := cw.WriteMajorTypeHeader(cbg.MajNegativeInt, uint64(-t.WinCount-1)); err != nil {
			return err
		}
	}

	// t.VRFProof (types.VRFPi) (slice)
	if len(t.VRFProof) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.VRFProof was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajByteString, uint64(len(t.VRFProof))); err != nil {
		return err
	}

	if _, err := cw.Write(t.VRFProof[:]); err != nil {
		return err
	}
	return nil
}

func (t *ElectionProof) UnmarshalCBOR(r io.Reader) (err error) {
	*t = ElectionProof{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.WinCount (int64) (int64)
	{
		maj, extra, err := cr.ReadHeader()
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative oveflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.WinCount = int64(extraI)
	}
	// t.VRFProof (types.VRFPi) (slice)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.VRFProof: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.VRFProof = make([]uint8, extra)
	}

	if _, err := io.ReadFull(cr, t.VRFProof[:]); err != nil {
		return err
	}
	return nil
}

var lengthBufBeaconEntry = []byte{130}

func (t *BeaconEntry) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufBeaconEntry); err != nil {
		return err
	}

	// t.Round (uint64) (uint64)

	if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.Round)); err != nil {
		return err
	}

	// t.Data ([]uint8) (slice)
	if len(t.Data) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.Data was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajByteString, uint64(len(t.Data))); err != nil {
		return err
	}

	if _, err := cw.Write(t.Data[:]); err != nil {
		return err
	}
	return nil
}

func (t *BeaconEntry) UnmarshalCBOR(r io.Reader) (err error) {
	*t = BeaconEntry{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Round (uint64) (uint64)

	{

		maj, extra, err = cr.ReadHeader()
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.Round = uint64(extra)

	}
	// t.Data ([]uint8) (slice)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.Data: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.Data = make([]uint8, extra)
	}

	if _, err := io.ReadFull(cr, t.Data[:]); err != nil {
		return err
	}
	return nil
}

var lengthBufMessageRoot = []byte{130}

func (t *MessageRoot) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufMessageRoot); err != nil {
		return err
	}

	// t.BlsRoot (cid.Cid) (struct)

	if err := cbg.WriteCid(cw, t.BlsRoot); err != nil {
		return xerrors.Errorf("failed to write cid field t.BlsRoot: %w", err)
	}

	// t.SecpkRoot (cid.Cid) (struct)

	if err := cbg.WriteCid(cw, t.SecpkRoot); err != nil {
		return xerrors.Errorf("failed to write cid field t.SecpkRoot: %w", err)
	}

	return nil
}

func (t *MessageRoot) UnmarshalCBOR(r io.Reader) (err error) {
	*t = MessageRoot{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.BlsRoot (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(cr)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.BlsRoot: %w", err)
		}

		t.BlsRoot = c

	}
	// t.SecpkRoot (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(cr)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.SecpkRoot: %w", err)
		}

		t.SecpkRoot = c

	}
	return nil
}

var lengthBufBlockMsg = []byte{131}

func (t *BlockMsg) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufBlockMsg); err != nil {
		return err
	}

	// t.Header (types.BlockHeader) (struct)
	if err := t.Header.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.BlsMessages ([]cid.Cid) (slice)
	if len(t.BlsMessages) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.BlsMessages was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajArray, uint64(len(t.BlsMessages))); err != nil {
		return err
	}
	for _, v := range t.BlsMessages {
		if err := cbg.WriteCid(w, v); err != nil {
			return xerrors.Errorf("failed writing cid field t.BlsMessages: %w", err)
		}
	}

	// t.SecpkMessages ([]cid.Cid) (slice)
	if len(t.SecpkMessages) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.SecpkMessages was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajArray, uint64(len(t.SecpkMessages))); err != nil {
		return err
	}
	for _, v := range t.SecpkMessages {
		if err := cbg.WriteCid(w, v); err != nil {
			return xerrors.Errorf("failed writing cid field t.SecpkMessages: %w", err)
		}
	}
	return nil
}

func (t *BlockMsg) UnmarshalCBOR(r io.Reader) (err error) {
	*t = BlockMsg{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 3 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Header (types.BlockHeader) (struct)

	{

		b, err := cr.ReadByte()
		if err != nil {
			return err
		}
		if b != cbg.CborNull[0] {
			if err := cr.UnreadByte(); err != nil {
				return err
			}
			t.Header = new(BlockHeader)
			if err := t.Header.UnmarshalCBOR(cr); err != nil {
				return xerrors.Errorf("unmarshaling t.Header pointer: %w", err)
			}
		}

	}
	// t.BlsMessages ([]cid.Cid) (slice)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.BlsMessages: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.BlsMessages = make([]cid.Cid, extra)
	}

	for i := 0; i < int(extra); i++ {

		c, err := cbg.ReadCid(cr)
		if err != nil {
			return xerrors.Errorf("reading cid field t.BlsMessages failed: %w", err)
		}
		t.BlsMessages[i] = c
	}

	// t.SecpkMessages ([]cid.Cid) (slice)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.SecpkMessages: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.SecpkMessages = make([]cid.Cid, extra)
	}

	for i := 0; i < int(extra); i++ {

		c, err := cbg.ReadCid(cr)
		if err != nil {
			return xerrors.Errorf("reading cid field t.SecpkMessages failed: %w", err)
		}
		t.SecpkMessages[i] = c
	}

	return nil
}

var lengthBufExpTipSet = []byte{131}

func (t *ExpTipSet) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufExpTipSet); err != nil {
		return err
	}

	// t.Cids ([]cid.Cid) (slice)
	if len(t.Cids) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.Cids was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajArray, uint64(len(t.Cids))); err != nil {
		return err
	}
	for _, v := range t.Cids {
		if err := cbg.WriteCid(w, v); err != nil {
			return xerrors.Errorf("failed writing cid field t.Cids: %w", err)
		}
	}

	// t.Blocks ([]*types.BlockHeader) (slice)
	if len(t.Blocks) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.Blocks was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajArray, uint64(len(t.Blocks))); err != nil {
		return err
	}
	for _, v := range t.Blocks {
		if err := v.MarshalCBOR(cw); err != nil {
			return err
		}
	}

	// t.Height (abi.ChainEpoch) (int64)
	if t.Height >= 0 {
		if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.Height)); err != nil {
			return err
		}
	} else {
		if err := cw.WriteMajorTypeHeader(cbg.MajNegativeInt, uint64(-t.Height-1)); err != nil {
			return err
		}
	}
	return nil
}

func (t *ExpTipSet) UnmarshalCBOR(r io.Reader) (err error) {
	*t = ExpTipSet{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 3 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Cids ([]cid.Cid) (slice)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.Cids: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.Cids = make([]cid.Cid, extra)
	}

	for i := 0; i < int(extra); i++ {

		c, err := cbg.ReadCid(cr)
		if err != nil {
			return xerrors.Errorf("reading cid field t.Cids failed: %w", err)
		}
		t.Cids[i] = c
	}

	// t.Blocks ([]*types.BlockHeader) (slice)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.Blocks: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.Blocks = make([]*BlockHeader, extra)
	}

	for i := 0; i < int(extra); i++ {

		var v BlockHeader
		if err := v.UnmarshalCBOR(cr); err != nil {
			return err
		}

		t.Blocks[i] = &v
	}

	// t.Height (abi.ChainEpoch) (int64)
	{
		maj, extra, err := cr.ReadHeader()
		var extraI int64
		if err != nil {
			return err
		}
		switch maj {
		case cbg.MajUnsignedInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 positive overflow")
			}
		case cbg.MajNegativeInt:
			extraI = int64(extra)
			if extraI < 0 {
				return fmt.Errorf("int64 negative oveflow")
			}
			extraI = -1 - extraI
		default:
			return fmt.Errorf("wrong type for int64 field: %d", maj)
		}

		t.Height = abi.ChainEpoch(extraI)
	}
	return nil
}

var lengthBufPaymentInfo = []byte{131}

func (t *PaymentInfo) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufPaymentInfo); err != nil {
		return err
	}

	// t.Channel (address.Address) (struct)
	if err := t.Channel.MarshalCBOR(cw); err != nil {
		return err
	}

	// t.WaitSentinel (cid.Cid) (struct)

	if err := cbg.WriteCid(cw, t.WaitSentinel); err != nil {
		return xerrors.Errorf("failed to write cid field t.WaitSentinel: %w", err)
	}

	// t.Vouchers ([]*paych.SignedVoucher) (slice)
	if len(t.Vouchers) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.Vouchers was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajArray, uint64(len(t.Vouchers))); err != nil {
		return err
	}
	for _, v := range t.Vouchers {
		if err := v.MarshalCBOR(cw); err != nil {
			return err
		}
	}
	return nil
}

func (t *PaymentInfo) UnmarshalCBOR(r io.Reader) (err error) {
	*t = PaymentInfo{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 3 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Channel (address.Address) (struct)

	{

		if err := t.Channel.UnmarshalCBOR(cr); err != nil {
			return xerrors.Errorf("unmarshaling t.Channel: %w", err)
		}

	}
	// t.WaitSentinel (cid.Cid) (struct)

	{

		c, err := cbg.ReadCid(cr)
		if err != nil {
			return xerrors.Errorf("failed to read cid field t.WaitSentinel: %w", err)
		}

		t.WaitSentinel = c

	}
	// t.Vouchers ([]*paych.SignedVoucher) (slice)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.Vouchers: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.Vouchers = make([]*paych.SignedVoucher, extra)
	}

	for i := 0; i < int(extra); i++ {

		var v paych.SignedVoucher
		if err := v.UnmarshalCBOR(cr); err != nil {
			return err
		}

		t.Vouchers[i] = &v
	}

	return nil
}

var lengthBufEvent = []byte{130}

func (t *Event) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufEvent); err != nil {
		return err
	}

	// t.Emitter (abi.ActorID) (uint64)

	if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.Emitter)); err != nil {
		return err
	}

	// t.Entries ([]types.EventEntry) (slice)
	if len(t.Entries) > cbg.MaxLength {
		return xerrors.Errorf("Slice value in field t.Entries was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajArray, uint64(len(t.Entries))); err != nil {
		return err
	}
	for _, v := range t.Entries {
		if err := v.MarshalCBOR(cw); err != nil {
			return err
		}
	}
	return nil
}

func (t *Event) UnmarshalCBOR(r io.Reader) (err error) {
	*t = Event{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 2 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Emitter (abi.ActorID) (uint64)

	{

		maj, extra, err = cr.ReadHeader()
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.Emitter = abi.ActorID(extra)

	}
	// t.Entries ([]types.EventEntry) (slice)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("t.Entries: array too large (%d)", extra)
	}

	if maj != cbg.MajArray {
		return fmt.Errorf("expected cbor array")
	}

	if extra > 0 {
		t.Entries = make([]EventEntry, extra)
	}

	for i := 0; i < int(extra); i++ {

		var v EventEntry
		if err := v.UnmarshalCBOR(cr); err != nil {
			return err
		}

		t.Entries[i] = v
	}

	return nil
}

var lengthBufEventEntry = []byte{132}

func (t *EventEntry) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)

	if _, err := cw.Write(lengthBufEventEntry); err != nil {
		return err
	}

	// t.Flags (uint8) (uint8)
	if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.Flags)); err != nil {
		return err
	}

	// t.Key (string) (string)
	if len(t.Key) > cbg.MaxLength {
		return xerrors.Errorf("Value in field t.Key was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len(t.Key))); err != nil {
		return err
	}
	if _, err := io.WriteString(w, string(t.Key)); err != nil {
		return err
	}

	// t.Codec (uint64) (uint64)

	if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(t.Codec)); err != nil {
		return err
	}

	// t.Value ([]uint8) (slice)
	if len(t.Value) > cbg.ByteArrayMaxLen {
		return xerrors.Errorf("Byte array in field t.Value was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajByteString, uint64(len(t.Value))); err != nil {
		return err
	}

	if _, err := cw.Write(t.Value[:]); err != nil {
		return err
	}
	return nil
}

func (t *EventEntry) UnmarshalCBOR(r io.Reader) (err error) {
	*t = EventEntry{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajArray {
		return fmt.Errorf("cbor input should be of type array")
	}

	if extra != 4 {
		return fmt.Errorf("cbor input had wrong number of fields")
	}

	// t.Flags (uint8) (uint8)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}
	if maj != cbg.MajUnsignedInt {
		return fmt.Errorf("wrong type for uint8 field")
	}
	if extra > math.MaxUint8 {
		return fmt.Errorf("integer in input was too large for uint8 field")
	}
	t.Flags = uint8(extra)
	// t.Key (string) (string)

	{
		sval, err := cbg.ReadString(cr)
		if err != nil {
			return err
		}

		t.Key = string(sval)
	}
	// t.Codec (uint64) (uint64)

	{

		maj, extra, err = cr.ReadHeader()
		if err != nil {
			return err
		}
		if maj != cbg.MajUnsignedInt {
			return fmt.Errorf("wrong type for uint64 field")
		}
		t.Codec = uint64(extra)

	}
	// t.Value ([]uint8) (slice)

	maj, extra, err = cr.ReadHeader()
	if err != nil {
		return err
	}

	if extra > cbg.ByteArrayMaxLen {
		return fmt.Errorf("t.Value: byte array too large (%d)", extra)
	}
	if maj != cbg.MajByteString {
		return fmt.Errorf("expected byte array")
	}

	if extra > 0 {
		t.Value = make([]uint8, extra)
	}

	if _, err := io.ReadFull(cr, t.Value[:]); err != nil {
		return err
	}
	return nil
}
