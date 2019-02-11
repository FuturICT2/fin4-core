module ExploreAssets.Model exposing (Model, init)

import Asset.Model exposing (Asset)
import Common.Json exposing (decodeAt)
import Common.Log exposing (log)
import ExploreAssets.Msg exposing (Msg(..))
import Material


type alias Model =
    { mdl : Material.Model
    , assetCategories : List String
    , activeCategory : String
    , assets : List Asset
    , favorites : List Asset
    , error : Maybe String
    }


init : Model
init =
    { mdl = Material.model
    , assetCategories = [ "Trends", "Sport", "Entertainment", "Tourism", "Business", "Lifestyle", "Nature" ]
    , activeCategory = "Trends"
    , assets = []
    , favorites = []
    , error = Nothing
    }
