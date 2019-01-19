module Person.Update exposing (update)

import Common.Json exposing (decodeAt, deocdeIntWithDefault)
import Common.Log exposing (log)
import Json.Decode as JD
import Main.Context exposing (Context)
import Material
import Person.Model exposing (Model)
import Person.Msg exposing (Msg(..))


update : Context -> Msg -> Model -> ( Model, Cmd Msg )
update ctx msg model =
    case msg of
        SelectTab tab ->
            { model | tab = tab } ! []

        OnPersonLoadResponse resp ->
            case resp of
                Ok data ->
                    { model
                        | data = Just data
                        , error = Nothing
                    }
                        ! []

                Err _ ->
                    { model
                        | data = Nothing
                        , error = Just "Error loading user page. Please try again."
                    }
                        ! []

        Mdl msg_ ->
            Material.update Mdl msg_ model
