module Main.Model exposing (Model, initModel)

import Asset.Model
import CreateAsset.Model
import ExploreAssets.Model
import Homepage.Model
import Main.Context exposing (Context, initContext)
import Main.Flags exposing (Flags)
import Main.Routing exposing (Route(..))
import Material
import Portfolio.Model
import Profile.Model
import UserLogin.Model


type alias Model =
    { v : String
    , context : Context
    , mdl : Material.Model
    , homepage : Homepage.Model.Model
    , exploreAssets : ExploreAssets.Model.Model
    , showMobileNav : Bool
    , portfolio : Portfolio.Model.Model
    , userlogin : UserLogin.Model.Model
    , createAsset : CreateAsset.Model.Model
    , asset : Asset.Model.Model
    , profile : Profile.Model.Model
    }


initModel : Flags -> Route -> Model
initModel flags route =
    { v = "1"
    , context = initContext flags route
    , mdl = Material.model
    , homepage = Homepage.Model.init
    , exploreAssets = ExploreAssets.Model.init
    , showMobileNav = False
    , portfolio = Portfolio.Model.init
    , userlogin = UserLogin.Model.init
    , createAsset = CreateAsset.Model.init
    , asset = Asset.Model.init
    , profile = Profile.Model.init
    }
