module CreateAsset.Model exposing (Model, init)

import CreateAsset.Types exposing (..)
import Http
import Material


type alias Model =
    { mdl : Material.Model
    , step : ActiveView
    , name : String
    , symbol : String
    , cap : String
    , decimals : String
    , purpose : String
    , isSensor : Bool
    , isMintable : Bool
    , isTransferable : Bool
    , isBurnable : Bool
    , isCreatingAsset : Bool
    , createAssetError : Maybe Http.Error
    , createdAssetId : Int
    , createdAssetName : String
    , createdAssetSymbol : String
    }


init : Model
init =
    { mdl = Material.model
    , step = FormView
    , name = ""
    , symbol = ""
    , cap = ""
    , decimals = ""
    , purpose = ""
    , isSensor = False
    , isMintable = False
    , isTransferable = False
    , isBurnable = False
    , isCreatingAsset = False
    , createAssetError = Nothing
    , createdAssetId = 0
    , createdAssetName = ""
    , createdAssetSymbol = ""
    }
