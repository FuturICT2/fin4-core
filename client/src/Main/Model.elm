module Main.Model exposing (Model, initModel)

import Actions.Model
import CreateToken.Model
import Main.Context exposing (Context, initContext)
import Main.Flags exposing (Flags)
import Main.Routing exposing (Route(..))
import Material
import Portfolio.Model
import Tokens.Model
import UserLogin.Model


type alias Model =
    { v : String
    , context : Context
    , mdl : Material.Model
    , showMobileNav : Bool
    , tokens : Tokens.Model.Model
    , portfolio : Portfolio.Model.Model
    , actions : Actions.Model.Model
    , createToken : CreateToken.Model.Model
    , userlogin : UserLogin.Model.Model
    }


initModel : Flags -> Route -> Model
initModel flags route =
    { v = "1"
    , context = initContext flags route
    , mdl = Material.model
    , showMobileNav = False
    , tokens = Tokens.Model.init
    , portfolio = Portfolio.Model.init
    , actions = Actions.Model.init
    , createToken = CreateToken.Model.init
    , userlogin = UserLogin.Model.init
    }
