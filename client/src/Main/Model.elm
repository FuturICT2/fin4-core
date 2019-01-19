module Main.Model exposing (Model, initModel)

import CreateToken.Model
import Main.Context exposing (Context, initContext)
import Main.Flags exposing (Flags)
import Main.Routing exposing (Route(..))
import Material
import Portfolio.Model
import Token.Model
import Tokens.Model
import UserLogin.Model


type alias Model =
    { v : String
    , context : Context
    , mdl : Material.Model
    , showMobileNav : Bool
    , token : Token.Model.Model
    , portfolio : Portfolio.Model.Model
    , tokens : Tokens.Model.Model
    , createToken : CreateToken.Model.Model
    , userlogin : UserLogin.Model.Model
    }


initModel : Flags -> Route -> Model
initModel flags route =
    { v = "1"
    , context = initContext flags route
    , mdl = Material.model
    , showMobileNav = False
    , token = Token.Model.init
    , portfolio = Portfolio.Model.init
    , tokens = Tokens.Model.init
    , createToken = CreateToken.Model.init
    , userlogin = UserLogin.Model.init
    }
