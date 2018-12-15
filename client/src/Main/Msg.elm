module Main.Msg exposing (Msg(..))

import Actions.Msg
import Common.Json exposing (EmptyResponse)
import CreateAction.Msg
import CreateToken.Msg
import Homepage.Homepage
import Http
import Main.User exposing (User)
import Material
import Navigation exposing (Location)
import Portfolio.Msg
import Tokens.Msg
import UserLogin.Msg
import Window


type Msg
    = OnRouteChange Location
    | ToggleMobileNav
    | OnCheckSessionResponse (Result Http.Error User)
    | Mdl (Material.Msg Msg)
    | OnWindowResize Window.Size
    | Homepage Homepage.Homepage.Msg
    | Tokens Tokens.Msg.Msg
    | Portfolio Portfolio.Msg.Msg
    | Actions Actions.Msg.Msg
    | CreateToken CreateToken.Msg.Msg
    | CreateAction CreateAction.Msg.Msg
    | UserLogin UserLogin.Msg.Msg
    | UserLogout
    | OnLogoutResponse (Result Http.Error EmptyResponse)
