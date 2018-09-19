module Main.ViewBody exposing (ifAuth, ifNotAuth, render)

import Homepage.Homepage
import Html exposing (..)
import Main.Model exposing (Model)
import Main.Msg exposing (Msg)
import Main.NotFound
import Main.Routing exposing (Route(..))
import Portfolio.View
import Tokens.View


render : Model -> Html Msg
render model =
    case model.context.route of
        HomepageRoute ->
            Homepage.Homepage.render model.context

        TokensRoute ->
            Tokens.View.render model.context model.tokens

        PortfolioRoute ->
            Portfolio.View.render model.context model.portfolio

        NotFoundRoute ->
            Main.NotFound.render model


ifAuth : Model -> Html Msg -> Html Msg
ifAuth model view =
    case model.context.user of
        Just _ ->
            view

        Nothing ->
            view


ifNotAuth : Model -> Html Msg -> Html Msg
ifNotAuth model view =
    case model.context.user of
        Just _ ->
            div [] []

        Nothing ->
            view
