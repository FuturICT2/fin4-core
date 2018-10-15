module Main.ViewBody exposing (render)

import CreateToken.View
import Homepage.Homepage
import Html exposing (..)
import Main.Model exposing (Model)
import Main.Msg exposing (Msg)
import Main.NotFound
import Main.Routing exposing (Route(..))
import Portfolio.View
import Tokens.View
import UserLogin.ViewLogin


render : Model -> Html Msg
render model =
    case model.context.route of
        HomepageRoute ->
            Homepage.Homepage.render model.context

        TokensRoute ->
            Html.map Main.Msg.Tokens <|
                Tokens.View.render model.context model.tokens

        PortfolioRoute ->
            ifAuth model <|
                Portfolio.View.render model.context model.portfolio

        CreateTokenRoute ->
            ifAuth model <|
                Html.map Main.Msg.CreateToken <|
                    CreateToken.View.render model.context model.createToken

        UserLoginRoute ->
            ifNotAuth model <| renderLogin model

        NotFoundRoute ->
            Main.NotFound.render model


renderLogin : Model -> Html Msg
renderLogin model =
    Html.map Main.Msg.UserLogin <| UserLogin.ViewLogin.render model.userlogin


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
