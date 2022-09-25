package topic

import (
	"fmt"
	"os"

	"github.com/dgraph-io/badger/v3"
)

type TopicService struct {
	storageLocationPrefix string
}

func New() (*TopicService, error) {
	service := &TopicService{
		storageLocationPrefix: "/tmp/sungoq",
	}

	return service, nil
}

func (service *TopicService) storage() (*badger.DB, error) {
	storageOpt := badger.DefaultOptions(
		fmt.Sprintf("%s/%s", service.storageLocationPrefix, "sungoq"),
	)
	store, err := badger.Open(
		storageOpt,
	)
	if err != nil {
		return nil, err
	}

	return store, nil
}

func (service *TopicService) GetAll() ([]string, error) {
	topics := make([]string, 0)
	serviceStorage, err := service.storage()
	if err != nil {
		return nil, err
	}
	defer serviceStorage.Close()

	err = serviceStorage.View(func(txn *badger.Txn) error {
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
	serviceStorage, err := service.storage()
	if err != nil {
		return err
	}
	defer serviceStorage.Close()

	storage, err := badger.Open(
		badger.DefaultOptions(
			fmt.Sprintf("%s/%s", service.storageLocationPrefix, name),
		),
	)

	if err != nil {
		return err
	}

	defer storage.Close()

	err = serviceStorage.Update(func(txn *badger.Txn) error {
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
	serviceStorage, err := service.storage()
	if err != nil {
		return err
	}
	defer serviceStorage.Close()

	err = os.RemoveAll(fmt.Sprintf("%s/%s", service.storageLocationPrefix, name))
	if err != nil {
		return err
	}

	err = serviceStorage.Update(func(txn *badger.Txn) error {
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
