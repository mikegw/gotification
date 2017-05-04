package notification

/*--- Public Data Types ---*/

type Model struct {
    Payload string `json:"payload"`
}


type SavedModel struct {
    Payload string `json:"payload"`
    Id string `json:"id"`
}



/*--- Public Functions ---*/

func Save(model Model) (SavedModel, error) {
    var saved SavedModel
    saved.Payload = model.Payload
    saved.Id = "some id"
    return saved, nil
}
