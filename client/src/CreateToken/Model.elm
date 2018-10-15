module CreateToken.Model exposing (Model, init)

import Common.Json exposing (decodeAt)
import Common.Log exposing (log)
import CreateToken.Msg exposing (Msg(..))
import Debug
import Http
import Json.Decode as JD
import Json.Decode.Pipeline as JP
import Material


type alias Model =
    { mdl : Material.Model
    , step : Int
    , name : String
    , symbol : String
    , description : String
    , hashtags : List String
    , newHashtag : String
    , shares : String
    , activeCategory : String
    , isCreatingToken : Bool
    , createTokenError : Maybe Http.Error
    }


init : Model
init =
    { mdl = Material.model
    , step = 0
    , name = ""
    , symbol = ""
    , description = ""
    , shares = ""
    , newHashtag = ""
    , hashtags = []
    , activeCategory = ""
    , isCreatingToken = False
    , createTokenError = Nothing
    }
