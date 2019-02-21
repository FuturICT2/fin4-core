module Main.ViewBody exposing (render)

import Asset.View
import CreateAsset.View
import ExploreAssets.View
import Homepage.View
import Html exposing (..)
import Main.Model exposing (Model)
import Main.Msg exposing (Msg(..))
import Main.NotFound
import Main.Routing exposing (Route(..))
import Portfolio.View
import Profile.View
import UserLogin.ViewFastSignup
import UserLogin.ViewForgot
import UserLogin.ViewForgotResetPass
import UserLogin.ViewLogin
import UserLogin.ViewSignup


render : Model -> Html Msg
render model =
    case model.context.route of
        HomepageRoute ->
            Html.map HomepageMsg <| Homepage.View.render model.context model.homepage

        ExploreAssetsRoute ->
            ifAuth model <|
                Html.map ExploreAssetsMsg <|
                    ExploreAssets.View.render model.context model.exploreAssets

        CreateAssetRoute ->
            ifAuth model <|
                Html.map CreateAssetMsg <|
                    CreateAsset.View.render model.context model.createAsset

        AssetRoute assetId ->
            ifAuth model <|
                Html.map AssetMsg <|
                    Asset.View.render model.context model.asset assetId

        ProfileRoute profileId ->
            ifAuth model <|
                Html.map ProfileMsg <|
                    Profile.View.render model.context model.profile profileId

        PortfolioRoute ->
            ifAuth model <|
                Portfolio.View.render model.context model.portfolio

        UserLoginRoute ->
            ifNotAuth model <| renderLogin model

        UserForgotRoute ->
            ifNotAuth model <| Html.map Main.Msg.UserLoginMsg <| UserLogin.ViewForgot.render model.userlogin

        UserForgotResetPassRoute _ _ ->
            Html.map Main.Msg.UserLoginMsg <| UserLogin.ViewForgotResetPass.render model.context model.userlogin

        UserSignupRoute ->
            ifNotAuth model <| Html.map Main.Msg.UserLoginMsg <| UserLogin.ViewSignup.render model.userlogin

        UserFastSignupRoute ->
            ifNotAuth model <| Html.map Main.Msg.UserLoginMsg <| UserLogin.ViewFastSignup.render model.userlogin

        NotFoundRoute ->
            Main.NotFound.render model


renderLogin : Model -> Html Msg
renderLogin model =
    Html.map Main.Msg.UserLoginMsg <| UserLogin.ViewLogin.render model.userlogin


ifAuth : Model -> Html Msg -> Html Msg
ifAuth model view =
    case model.context.user of
        Just _ ->
            view

        Nothing ->
            renderLogin model


ifNotAuth : Model -> Html Msg -> Html Msg
ifNotAuth model view =
    case model.context.user of
        Just _ ->
            div [] []

        Nothing ->
            view
