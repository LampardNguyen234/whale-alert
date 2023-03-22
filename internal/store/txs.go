package store

// StoreProcessedTx marks the given txHash as processed.
func (s *Store) StoreProcessedTx(txHash []byte) error {
	return s.db.SetByKey(makeKey(processedTxKey, txHash...), []byte{1})
}

// IsTxProcessed checks if the given txHash has been processed.
func (s *Store) IsTxProcessed(txHash []byte) (bool, error) {
	data, err := s.db.GetByKey(makeKey(processedTxKey, txHash...))
	if err != nil {
		return false, err
	}
	if len(data) == 0 {
		return false, nil
	}

	return data[0] == 1, nil
}
