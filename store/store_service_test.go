import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testStoreService = &StorageService{}

func init() {
	testStoreService = InitializeStore()
}

