module CreateAsset.Model exposing (Model, init)

import CreateAsset.Types exposing (..)
import Http
import Material


type alias Model =
    { mdl : Material.Model
    , step : ActiveView
    , name : String
    , symbol : String
    , purpose : String
    , isCreatingAsset : Bool
    , createAssetError : Maybe Http.Error
    , createdAssetId : Int
    , createdAssetName : String
    , createdAssetSymbol : String
    }


init : Model
init =
    { mdl = Material.model
    , step = InformationView
    , name = ""
    , symbol = ""
    , purpose = ""
    , isCreatingAsset = False
    , createAssetError = Nothing
    , createdAssetId = 0
    , createdAssetName = ""
    , createdAssetSymbol = ""
    }
