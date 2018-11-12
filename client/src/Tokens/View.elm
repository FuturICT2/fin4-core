module Tokens.View exposing (render, renderData, renderRow)

import Common.Decimal exposing (renderDecimal)
import Common.Error as Error
import Common.Styles exposing (padding, textLeft, textRight, toMdlCss)
import Html exposing (..)
import Html.Attributes exposing (href, style)
import List exposing (reverse, sortBy)
import Main.Context exposing (Context)
import Main.User exposing (User)
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
import Random
import Tokens.Model exposing (Model)
import Tokens.Msg exposing (Msg(..))


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
    div [ style [ ( "padding-top", "15px" ), ( "width", "100%" ) ] ]
        [ Tabs.render Mdl
            [ 0 ]
            model.mdl
            [ Tabs.ripple
            , Tabs.activeTab model.selectedTab
            , Tabs.onSelectTab SelectTab
            ]
            [ Tabs.label
                [ Options.center ]
                [ Icon.view "trending_up" [ Icon.size18, css "color" "red" ]
                , Options.span [ css "width" "4px" ] []
                , text "Tokens"
                ]
            , Tabs.label
                [ Options.center ]
                [ Icon.view "people" [ Icon.size18, css "color" "red" ]
                , Options.span [ css "width" "4px" ] []
                , text "people"
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
                                    div [] [ text "We are not able to fetch tokens. Please try again" ]

                                Just tokens ->
                                    renderData ctx model tokens

                _ ->
                    case model.error of
                        Just _ ->
                            Error.renderMaybeError model.error

                        Nothing ->
                            case model.tokens of
                                Nothing ->
                                    div [] [ text "..." ]

                                Just tokens ->
                                    div [] <| List.map (renderUser model) <| reverse tokens.people
            ]
        ]


renderUser : Model -> User -> Html Msg
renderUser model user =
    Card.view
        [ css "border" "1px solid #ddd"
        , css "width" "100%"
        , css "margin-top" "15px"
        ]
        [ Card.title []
            [ Card.head []
                [ Lists.content []
                    [ text user.name ]
                ]
            ]
        ]


renderData : Context -> Model -> Tokens -> Html Msg
renderData ctx model tokens =
    let
        sorted =
            reverse tokens.entries

        userId =
            case ctx.user of
                Just user ->
                    user.id

                Nothing ->
                    0
    in
    case List.length sorted > 0 of
        False ->
            div
                [ style
                    [ ( "margin", "auto" )
                    , ( "padding-top", "60px" )
                    , ( "text-align", "center" )
                    ]
                ]
                [ Options.styled p
                    [ Typo.caption
                    ]
                    [ text "No tokens has been created yet!" ]
                ]

        True ->
            div [] <| List.map (renderRow model) sorted


renderRow : Model -> Token -> Html Msg
renderRow model token =
    -- let
    --     gen =
    --         Random.float 0 4
    -- in
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
                    , small [] [ text <| " (" ++ token.symbol ++ ")" ]
                    ]
                ]
            ]
        , Card.text
            []
            [ text token.purpose ]
        , Card.actions
            [ Card.border, css "vertical-align" "center" ]
            [ Button.render Mdl
                [ 0 ]
                model.mdl
                [ Options.onClick (DoLike token.id) ]
                [ Icon.i "favorite_border"
                , text <| " " ++ toString token.favouriteCount
                ]

            -- , Button.render Mdl
            --     [ 1, 0 ]
            --     model.mdl
            --     [ Button.ripple
            --     , Button.accent
            --     , Button.link <| "https://rinkeby.etherscan.io/tx/" ++ token.txAddress
            --
            --     -- , Options.attribute <| Html.Attributes.target "_blank"
            --     , css "float" "right"
            --     ]
            --     [ text "transaction" ]
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
