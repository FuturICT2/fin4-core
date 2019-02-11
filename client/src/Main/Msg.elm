module Main.Msg exposing (Msg(..))

import Asset.Msg
import Common.Json exposing (EmptyResponse)
import CreateAsset.Msg
import CreateToken.Msg
import ExploreAssets.Msg
import Homepage.Homepage
import Homepage.Msg
import Http
import Main.User exposing (User)
import Material
import Navigation exposing (Location)
import Person.Msg
import Portfolio.Msg
import Profile.Msg
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
    | HomepageMsg Homepage.Msg.Msg
    | ExploreAssetsMsg ExploreAssets.Msg.Msg
    | Token Token.Msg.Msg
    | Person Person.Msg.Msg
    | Portfolio Portfolio.Msg.Msg
    | Tokens Tokens.Msg.Msg
    | CreateToken CreateToken.Msg.Msg
    | UserLoginMsg UserLogin.Msg.Msg
    | UserLogout
    | OnLogoutResponse (Result Http.Error EmptyResponse)
    | CreateAssetMsg CreateAsset.Msg.Msg
    | AssetMsg Asset.Msg.Msg
    | ProfileMsg Profile.Msg.Msg
