package notification

import (
    "fmt"
    "io"
    "io/ioutil"
    "encoding/json"
    "net/http"
)



/*--- Public Functions ---*/

func Create(writer http.ResponseWriter, request *http.Request, persistor Persistor) (int, error) {
    model, err := fetchModel(request)
    if badRequest(err) {
        writeBadRequest(err, writer)
        return 400, nil
    }
    if err != nil {
        return 500, err
    }

    savedModel, err := Save(model, persistor)
    if badRequest(err) {
        writeBadRequest(err, writer)
        return 400, nil
    }
    if err != nil {
        return 500, err
    }

    modelResponse, err := buildResponse(savedModel)
    if badRequest(err) {
        writeBadRequest(err, writer)
        return 400, nil
    }
    if err != nil {
        return 500, err
    }

    writeResponse(modelResponse, writer)
    return 200, nil
}


/*--- Private Data Types ---*/
type badRequestError struct {
    Message string
}

func (err badRequestError) Error() string {
    return err.Message
}


/*--- Private Functions ---*/

func fetchModel(request *http.Request) (Model, error) {
    var model Model

    rawBody, err := readBody(request)
    if err != nil {
        return model, err
    }

    err = json.Unmarshal(rawBody, &model)
    if err != nil {
        return model, badRequestError{"Invalid JSON"}
    }

    err = validateModel(model)
    if err != nil {
        return model, err
    }

    return model, nil
}

func buildResponse(model SavedModel) (string, error) {
    var response string
    responseBytes, err := json.Marshal(model)
    if err != nil {
        return response, err
    }
    response = string(responseBytes)
    return response, nil
}

func readBody(request *http.Request) ([]byte, error) {
    return ioutil.ReadAll(io.LimitReader(request.Body, 256))
}

func writeResponse(modelResponse string, writer http.ResponseWriter) {
    fmt.Fprintf(writer, modelResponse)
}

func validateModel(model Model) error {
    if model.Payload == "" {
        return badRequestError{"Missing JSON parameter: payload"}
    }
    return nil
}

func badRequest(err error) bool {
    _, badRequest := err.(badRequestError)
    return badRequest
}

func writeBadRequest(err error, writer http.ResponseWriter) {
    errorMessage := fmt.Sprintf("{\"message\":\"%s\"}", err.Error())
    http.Error(writer, errorMessage, 400)
}
