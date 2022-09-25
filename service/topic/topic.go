package topic

import (
	"fmt"
	"os"

	"github.com/dgraph-io/badger/v3"
)

type TopicService struct {
	storageLocationPrefix string
	storage               *badger.DB
}

func New() (*TopicService, error) {
	service := &TopicService{
		storageLocationPrefix: "/tmp/sungoq",
	}

	storageOpt := badger.DefaultOptions(
		fmt.Sprintf("%s/%s", service.storageLocationPrefix, "sungoq"),
	)
	storage, err := badger.Open(
		storageOpt,
	)
	if err != nil {
		return nil, err
	}

	service.storage = storage
	return service, nil
}

func (service *TopicService) GetAll() ([]string, error) {
	topics := make([]string, 0)

	err := service.storage.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			k := item.Key()
			topics = append(topics, string(k))
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return topics, nil
}

func (service *TopicService) Create(name string) error {
	_, err := badger.Open(
		badger.DefaultOptions(
			fmt.Sprintf("%s/%s", service.storageLocationPrefix, name),
		),
	)

	if err != nil {
		return err
	}

	err = service.storage.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(name), []byte(name))
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (service *TopicService) Delete(name string) error {
	err := os.RemoveAll(fmt.Sprintf("%s/%s", service.storageLocationPrefix, name))
	if err != nil {
		return err
	}

	err = service.storage.Update(func(txn *badger.Txn) error {
		err := txn.Delete([]byte(name))
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
