module Main.Model exposing (Model, initModel)

import Asset.Model
import CreateAsset.Model
import CreateToken.Model
import Homepage.Model
import Main.Context exposing (Context, initContext)
import Main.Flags exposing (Flags)
import Main.Routing exposing (Route(..))
import Material
import Person.Model
import Portfolio.Model
import Profile.Model
import Token.Model
import Tokens.Model
import UserLogin.Model


type alias Model =
    { v : String
    , context : Context
    , mdl : Material.Model
    , homepage : Homepage.Model.Model
    , showMobileNav : Bool
    , token : Token.Model.Model
    , portfolio : Portfolio.Model.Model
    , tokens : Tokens.Model.Model
    , createToken : CreateToken.Model.Model
    , userlogin : UserLogin.Model.Model
    , person : Person.Model.Model
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
    , showMobileNav = False
    , token = Token.Model.init
    , portfolio = Portfolio.Model.init
    , tokens = Tokens.Model.init
    , createToken = CreateToken.Model.init
    , userlogin = UserLogin.Model.init
    , person = Person.Model.init
    , createAsset = CreateAsset.Model.init
    , asset = Asset.Model.init
    , profile = Profile.Model.init
    }
