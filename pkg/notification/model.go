package notification

// import "github.com/aws/aws-sdk-go/service/sqs"

/*--- Public Data Types ---*/

type Model struct {
    Payload string `json:"payload"`
}


type SavedModel struct {
    Payload string `json:"payload"`
    ID string `json:"id"`
}

type Persistor interface {
    Persist(Model) (persistedID string, err error)
}



/*--- Public Functions ---*/

func Save(model Model, persistor Persistor) (SavedModel, error) {
    var saved SavedModel
    ID, err := persistor.Persist(model)
    if err != nil {
        return saved, err
    }
    saved.Payload = model.Payload
    saved.ID = ID
    return saved, nil
}
