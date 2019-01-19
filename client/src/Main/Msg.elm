module Main.Msg exposing (Msg(..))

import Common.Json exposing (EmptyResponse)
import CreateToken.Msg
import Homepage.Homepage
import Http
import Main.User exposing (User)
import Material
import Navigation exposing (Location)
import Person.Msg
import Portfolio.Msg
import Token.Msg
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
    | Token Token.Msg.Msg
    | Person Person.Msg.Msg
    | Portfolio Portfolio.Msg.Msg
    | Tokens Tokens.Msg.Msg
    | CreateToken CreateToken.Msg.Msg
    | UserLoginMsg UserLogin.Msg.Msg
    | UserLogout
    | OnLogoutResponse (Result Http.Error EmptyResponse)
