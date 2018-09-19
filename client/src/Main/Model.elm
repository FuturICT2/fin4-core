module Main.Model exposing (Model, initModel)

import Main.Context exposing (Context, initContext)
import Main.Flags exposing (Flags)
import Main.Routing exposing (Route(..))
import Material
import Portfolio.Model
import Tokens.Model


type alias Model =
    { v : String
    , context : Context
    , mdl : Material.Model
    , showMobileNav : Bool
    , tokens : Tokens.Model.Model
    , portfolio : Portfolio.Model.Model
    }


initModel : Flags -> Route -> Model
initModel flags route =
    { v = "1"
    , context = initContext flags route
    , mdl = Material.model
    , showMobileNav = False
    , tokens = Tokens.Model.init
    , portfolio = Portfolio.Model.init
    }
