module Tokens.View exposing (render, renderData, renderRow)

import Common.Decimal exposing (renderDecimal)
import Common.Error as Error
import Common.Styles exposing (padding, textLeft, textRight, toMdlCss)
import Html exposing (..)
import Html.Attributes exposing (style)
import Main.Context exposing (Context)
import Main.Msg exposing (Msg(..))
import Material.Button as Button
import Material.Card as Card
import Material.Color as Color
import Material.Icon as Icon
import Material.List as Lists
import Material.Options as Options exposing (css)
import Material.Table as Table
import Material.Tabs as Tabs
import Material.Typography as Typo
import Model.Tokens exposing (Token, Tokens)
import Tokens.Model exposing (Model)


render : Context -> Model -> Html Msg
render ctx model =
    let
        value =
            case model.tokens of
                Just tokens ->
                    tokens.valueInUSD

                Nothing ->
                    "0.00"
    in
    div [ style [ ( "padding-top", "15px" ) ] ]
        [ Options.styled p [ Typo.headline ] [ text "Tokens" ]
        , case model.error of
            Just _ ->
                Error.renderMaybeError model.error

            Nothing ->
                case model.tokens of
                    Nothing ->
                        div [] [ text "..." ]

                    Just tokens ->
                        renderData ctx model tokens
        ]


renderData : Context -> Model -> Tokens -> Html Msg
renderData ctx model tokens =
    case List.length tokens.entries > 0 of
        False ->
            div []
                [ Options.styled p
                    [ Typo.caption ]
                    [ text "empty" ]
                ]

        True ->
            div []
                [ div []
                    [ text <|
                        toString tokens.count
                            ++ " listed tokens of value "
                            ++ tokens.valueInUSD
                            ++ " USD"
                    ]
                , Table.table []
                    [ Table.thead []
                        [ Table.tr []
                            [ Table.th [ toMdlCss textLeft ] [ text "Token" ]
                            , Table.th [ toMdlCss textRight ] [ text "Supply" ]
                            , Table.th [ toMdlCss textRight ] [ text "24 Change" ]
                            ]
                        ]
                    , Table.tbody [] <| List.map renderRow tokens.entries
                    ]
                ]


renderRow : Token -> Html Msg
renderRow token =
    Table.tr []
        [ Table.td [ toMdlCss textLeft ] [ text token.name ]
        , Table.td [ toMdlCss textRight ] [ renderDecimal token.totalSupply, text token.symbol ]
        , Table.td [ toMdlCss textRight ] [ renderChange token.change24 ]
        ]


renderChange : Float -> Html Msg
renderChange value =
    let
        color =
            case value > 0 of
                True ->
                    "green"

                False ->
                    "red"

        sign =
            case value > 0 of
                True ->
                    "+"

                False ->
                    ""
    in
    span [ style [ ( "color", color ) ] ]
        [ text <| sign ++ toString value ++ "%" ]
