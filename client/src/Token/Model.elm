module Token.Model exposing (Model, init)

import Array exposing (Array)
import Http
import Material
import Model.Tokens exposing (Token)
import Token.Msg exposing (Msg(..))


type alias Image =
    { contents : String
    , filename : String
    }


type alias Model =
    { mdl : Material.Model
    , data : Maybe Token
    , error : Maybe String
    , img : Maybe Image
    , proposalText : String
    , tab : Int
    , showClaimForm : Bool
    , submitProposalError : Maybe Http.Error
    }


init : Model
init =
    { mdl = Material.model
    , data = Nothing
    , error = Nothing
    , img = Nothing
    , proposalText = ""
    , tab = 0
    , showClaimForm = False
    , submitProposalError = Nothing
    }
