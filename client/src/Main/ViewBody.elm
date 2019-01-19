module Main.ViewBody exposing (render)

import Actions.View
import CreateAction.View
import CreateToken.View
import Homepage.Homepage
import Html exposing (..)
import Main.Model exposing (Model)
import Main.Msg exposing (Msg)
import Main.NotFound
import Main.Routing exposing (Route(..))
import Portfolio.View
import Token.View
import Tokens.View
import UserLogin.ViewFastSignup
import UserLogin.ViewForgot
import UserLogin.ViewForgotResetPass
import UserLogin.ViewLogin
import UserLogin.ViewSignup


render : Model -> Html Msg
render model =
    case model.context.route of
        HomepageRoute ->
            ifAuth model <|
                Html.map Main.Msg.Tokens <|
                    Tokens.View.render model.context model.tokens

        TokensRoute ->
            ifAuth model <|
                Html.map Main.Msg.Tokens <|
                    Tokens.View.render model.context model.tokens

        TokenRoute tokenId ->
            ifAuth model <|
                Html.map Main.Msg.Token <|
                    Token.View.render model.context model.token tokenId

        PortfolioRoute ->
            ifAuth model <|
                Portfolio.View.render model.context model.portfolio

        ActionsRoute ->
            ifAuth model <|
                Html.map Main.Msg.Actions <|
                    Actions.View.render model.context model.actions

        CreateTokenRoute ->
            ifAuth model <|
                Html.map Main.Msg.CreateToken <|
                    CreateToken.View.render model.context model.createToken

        CreateActionRoute ->
            ifAuth model <|
                Html.map Main.Msg.CreateAction <|
                    CreateAction.View.render model.context model.createAction

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
