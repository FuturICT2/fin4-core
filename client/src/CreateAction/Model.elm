module CreateAction.Model exposing (Model, init)

import Common.Json exposing (decodeAt)
import Common.Log exposing (log)
import CreateAction.Msg exposing (Msg(..))
import Debug
import Http
import Json.Decode as JD
import Json.Decode.Pipeline as JP
import Material


type alias Model =
    { mdl : Material.Model
    , step : Int
    , timeLimit : String
    , description : String
    , isCreatingAction : Bool
    , createActionError : Maybe Http.Error
    }


init : Model
init =
    { mdl = Material.model
    , step = 0
    , timeLimit = ""
    , description = ""
    , isCreatingAction = False
    , createActionError = Nothing
    }
