module Portfolio.View exposing (render)

import Common.Decimal exposing (renderDecimal)
import Common.Error as Error
import Common.Styles exposing (padding, textLeft, textRight, toMdlCss)
import Html exposing (..)
import Html.Attributes exposing (style)
import Html.Events exposing (onClick)
import Main.Context exposing (Context)
import Main.Msg exposing (Msg(..))
import Material.Options as Options
import Material.Table as Table
import Material.Typography as Typo
import Model.Portfolio exposing (Portfolio, Position)
import Portfolio.Model exposing (Model)


render : Context -> Model -> Html Msg
render ctx model =
    let
        portfolioValue =
            case model.portfolio of
                Just portfolio ->
                    portfolio.valueInUSD

                Nothing ->
                    "0.00"

        userName =
            case ctx.user of
                Just user ->
                    user.name

                Nothing ->
                    ""
    in
    div []
        [ Options.styled p [ Typo.headline ] [ text "Profile" ]
        , text <| "Welcome, " ++ userName ++ "! ("
        , a [ onClick Main.Msg.UserLogout ] [ text "logout" ]
        , text ")"
        , case model.error of
            Just _ ->
                Error.renderMaybeError model.error

            Nothing ->
                case model.portfolio of
                    Nothing ->
                        div [] [ text "..." ]

                    Just portfolio ->
                        renderData ctx model portfolio
        ]


renderData : Context -> Model -> Portfolio -> Html Msg
renderData ctx model portfolio =
    case List.length portfolio.positions > 0 of
        False ->
            div []
                [ Options.styled p
                    [ Typo.caption ]
                    [ text "" ]
                ]

        True ->
            div []
                []
