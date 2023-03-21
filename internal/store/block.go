package store

import (
	"fmt"
	"math/big"
)

// StoreLastBlock stores the latest block for a given chainID to the db.
func (s *Store) StoreLastBlock(blk *big.Int, chainID uint8) error {
	if blk == nil {
		return fmt.Errorf("nil blk")
	}

	return s.db.SetByKey(makeKey(lastBlkKey, chainID), blk.Bytes())
}

// GetLastBlock retrieves the latest block for a given chainID from the db.
func (s *Store) GetLastBlock(chainID uint8) (*big.Int, error) {
	data, err := s.db.GetByKey(makeKey(lastBlkKey, chainID))
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, fmt.Errorf("empty data")
	}

	blk := new(big.Int).SetBytes(data)
	return blk, nil
}
