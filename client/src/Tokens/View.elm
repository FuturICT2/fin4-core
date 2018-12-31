module Tokens.View exposing (render, renderData, renderRow)

import Common.Decimal exposing (renderDecimal)
import Common.Error as Error
import Common.Styles exposing (padding, textLeft, textRight, toMdlCss)
import Html exposing (..)
import Html.Attributes exposing (href, src, style)
import Html.Events exposing (onClick)
import Identicon exposing (identicon)
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
    div [ style [ ( "padding-top", "15px" ), ( "width", "100%" ) ] ]
        [ header [ style [ ( "text-align", "center" ) ] ] [ text "Demo tokens" ]
        , case model.error of
            Just _ ->
                Error.renderMaybeError model.error

            Nothing ->
                case model.tokens of
                    Nothing ->
                        div [] [ text "We are not able to fetch tokens. Please try again" ]

                    Just tokens ->
                        renderData ctx model tokens
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
            tokens.entries

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
    let
        likeBackground =
            case token.didUserLike of
                True ->
                    "red"

                False ->
                    "inherit"
    in
    div
        [ style
            [ ( "border", "1px solid #ddd" )
            , ( "margin-bottom", "30px" )
            , ( "border-radius", "8px" )
            ]
        ]
        [ div
            [ style
                [ ( "display", "flex" )
                , ( "padding", "5px 0 0 5px" )
                ]
            ]
            [ div
                [ style
                    [ ( "padding", "15px" )
                    , ( "text-align", "center" )
                    , ( "width", "60px" )
                    , ( "height", "60px" )
                    , ( "border", "1px solid #ddd" )
                    , ( "border-radius", "50%" )
                    ]
                ]
                [ identicon "30px" token.name
                ]
            , div
                [ style
                    [ ( "padding", "8px" )
                    ]
                ]
                [ header
                    [ style
                        [ ( "margin", "0 0 10px 0" )
                        ]
                    ]
                    [ text token.name
                    ]
                , div
                    [ style
                        [ ( "margin-bottom", "15px" ) ]
                    ]
                    [ text <| "total supply: " ++ token.totalSupply
                    ]
                ]
            ]
        , div
            [ actionButtonsStyle
            ]
            [ Button.render Mdl
                [ token.id ]
                model.mdl
                [ Button.raised
                , Button.ripple
                , Options.onClick (DoLike token.id token.didUserLike)
                , toMdlCss buttonStyle
                ]
                [ Icon.view "favorite" [ Icon.size24, css "color" likeBackground ]
                , text <| " " ++ toString token.favouriteCount
                ]
            ]
        ]


tokenStyle =
    style
        [ ( "border", "1px solid #ddd" )
        , ( "margin-top", "15px" )
        , ( "padding", "5px" )
        ]


actionButtonsStyle =
    style
        [ ( "border-top", "1px solid #ddd" )
        , ( "background", "#f2f2f2" )
        , ( "font-weight", "bold" )
        , ( "display", "block" )
        ]


buttonStyle =
    style
        [ ( "width", "100%" )
        , ( "text-align", "center" )
        , ( "line-height", "normal" )
        , ( "text-decoration", "none" )
        , ( "color", "inherit" )
        , ( "display", "inline-block" )
        , ( "box-shadow", "none" )
        , ( "font-size", "20px" )
        ]
