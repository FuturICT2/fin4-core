module Tokens.View exposing (render, renderData, renderRow)

import Common.Decimal exposing (renderDecimal)
import Common.Error as Error
import Common.Styles exposing (padding, textLeft, textRight, toMdlCss)
import Html exposing (..)
import Html.Attributes exposing (style)
import List exposing (reverse, sortBy)
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
import Tokens.Msg


white : Options.Property c m
white =
    Color.text Color.white


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
        [ Tabs.render Mdl
            [ 0 ]
            model.mdl
            [ Tabs.ripple
            , Tabs.activeTab model.selectedTab
            ]
            [ Tabs.label
                [ Options.center ]
                [ Options.span [ css "width" "4px" ] []
                , text "trending"
                ]
            , Tabs.label
                [ Options.center ]
                [ Options.span [ css "width" "4px" ] []
                , text "health"
                ]
            , Tabs.label
                [ Options.center ]
                [ Options.span [ css "width" "4px" ] []
                , text "climate"
                ]
            ]
            [ case model.selectedTab of
                0 ->
                    case model.error of
                        Just _ ->
                            Error.renderMaybeError model.error

                        Nothing ->
                            case model.tokens of
                                Nothing ->
                                    div [] [ text "..." ]

                                Just tokens ->
                                    renderData ctx model tokens

                _ ->
                    div [] [ text "tab2" ]
            ]
        ]


renderData : Context -> Model -> Tokens -> Html Msg
renderData ctx model tokens =
    let
        sorted =
            reverse <| sortBy .change24 tokens.entries
    in
    case List.length sorted > 0 of
        False ->
            div []
                [ Options.styled p
                    [ Typo.caption ]
                    [ text "empty" ]
                ]

        True ->
            div [] <| List.map (renderRow model) tokens.entries


renderRow : Model -> Token -> Html Msg
renderRow model token =
    Card.view
        [ css "border" "1px solid #ddd"
        , css "width" "100%"
        , css "margin-top" "15px"
        ]
        [ Card.title []
            [ Card.head []
                [ Lists.content []
                    [ Lists.avatarImage token.logo []
                    , text <| " " ++ token.name
                    ]
                ]
            ]
        , Card.text [] [ text token.symbol ]
        , Card.actions
            [ Card.border, css "vertical-align" "center" ]
            [ Button.render Mdl
                [ 0 ]
                model.mdl
                [ Button.colored
                ]
                [ text "buy" ]
            , Button.render Mdl
                [ 0 ]
                model.mdl
                [ Button.colored
                ]
                [ text "sell" ]
            , Button.render Mdl
                [ 8, 1 ]
                model.mdl
                [ css "padding" "0", css "float" "right" ]
                [ renderChange token.change24 ]
            ]
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

        icon =
            case value > 0 of
                True ->
                    Icon.i "arrow_drop_up"

                False ->
                    Icon.i "arrow_drop_down"
    in
    span [ change24Style color ]
        [ icon, text <| toString (abs value) ++ "%" ]


change24Style color =
    style
        [ ( "color", color )
        , ( "float", "right" )
        ]


tokenStyle =
    style
        [ ( "border", "1px solid #ddd" )
        , ( "margin-top", "15px" )
        , ( "padding", "5px" )
        ]
