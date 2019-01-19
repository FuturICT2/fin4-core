module Tokens.View exposing (render)

import Common.Decimal exposing (renderDecimal, renderDecimalWithPrecision)
import Common.Error as Error
import Common.Styles exposing (padding, textLeft, textRight, toMdlCss)
import Dict exposing (Dict)
import Html exposing (..)
import Html.Attributes exposing (id, placeholder, rows, src, style, type_, value)
import Html.Events exposing (on, onClick, onInput)
import Identicon exposing (identicon)
import Json.Decode as JD
import Main.Context exposing (Context)
import Material.Button as Button
import Material.Chip as Chip
import Material.Icon as Icon
import Material.Options as Options exposing (css)
import Material.Table as Table
import Material.Tooltip as Tooltip
import Material.Typography as Typo
import Model.Tokens exposing (Token, Tokens)
import Tokens.Model exposing (Model)
import Tokens.Msg exposing (Msg(..))


render : Context -> Model -> Html Msg
render ctx model =
    let
        content =
            case model.tokens of
                Just tokens ->
                    case List.length tokens.entries of
                        0 ->
                            renderEmpty

                        _ ->
                            renderTokens ctx model tokens

                Nothing ->
                    renderEmpty
    in
    content


renderTokens : Context -> Model -> Tokens -> Html Msg
renderTokens ctx model tokens =
    div
        []
        [ div
            [ style
                [ ( "margin", "30px 0" )
                , ( "text-align", "center" )
                ]
            ]
            [ text "stats: "
            , a [] [ text "344 users " ]
            , text " | "
            , a [] [ text "233 oracles" ]
            , text " | "
            , a [] [ text "33 miners" ]
            , br [] []
            , text "Finfour tokens are socially mined. Click mine to participate!"
            ]
        , div [] <| List.map (renderToken ctx model) tokens.entries
        ]


renderToken ctx model token =
    div
        [ style
            [ ( "border", "1px solid #ddd" )
            , ( "border-radius", "8px" )
            , ( "margin-bottom", "15px" )
            ]
        ]
        [ renderTokenInfo ctx model token

        -- , renderTokenClaims ctx model token
        -- , renderClaimForm ctx model token
        -- , renderTokenControls ctx model token
        ]


renderTokenInfo ctx model token =
    let
        likeBackground =
            case token.didUserLike of
                True ->
                    "red"

                False ->
                    "inherit"
    in
    div
        []
        [ div
            [ style
                [ ( "display", "flex" )
                , ( "padding", "5px 0 0 5px" )
                ]
            ]
            [ div
                [ style
                    [ ( "padding", "10px" )
                    , ( "text-align", "center" )
                    , ( "width", "40px" )
                    , ( "height", "40px" )
                    , ( "border", "1px solid #ddd" )
                    , ( "border-radius", "50%" )
                    ]
                ]
                [ identicon "20px" token.name
                ]
            , div []
                [ header
                    [ style
                        [ ( "margin", "0 0 10px 0" )
                        , ( "padding", "10px" )
                        , ( "font-size", "26px" )
                        ]
                    ]
                    [ text token.name
                    ]
                ]
            ]
        , div
            [ style
                [ ( "padding", "7px" )
                , ( "line-height", "1" )
                , ( "font-size", "18px" )
                , ( "border-bottom", "1px solid #ddd" )
                ]
            ]
            [ text token.purpose
            ]
        , div
            [ style
                [ ( "padding", "15px 5px" )
                ]
            ]
            [ b []
                [ text <|
                    toString
                        (List.length <|
                            List.filter (\v -> v.isApproved == True) token.claims
                        )
                , small [] [ text <| "" ++ token.symbol ]
                ]
            , text " have been mined"
            , div [ style [ ( "float", "right" ) ] ]
                [ text "oracle: "
                , a [] [ text token.creatorName ]
                ]
            ]
        , renderTokenControls ctx model token
        ]


renderTokenControls ctx model token =
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
            , ( "border-top", "none" )
            , ( "border-radius", "0 0 8px 8px" )
            , ( "background", "#eee" )
            ]
        ]
        [ div [ cardBtnsStyle ]
            [ div
                [ cardBtnStyle
                ]
                [ Button.render Mdl
                    [ token.id ]
                    model.mdl
                    [ Button.raised
                    , Button.ripple
                    , Options.onClick (DoLike token.id token.didUserLike)
                    , toMdlCss buttonStyle
                    , css "width" "25%"
                    , css "border-right" "1px solid rgb(199, 199, 199)"
                    ]
                    [ Icon.view "favorite" [ Icon.size24, css "color" likeBackground ]
                    , text <| " " ++ toString token.favouriteCount
                    ]
                , Button.render Mdl
                    [ token.id, 2 ]
                    model.mdl
                    [ Button.raised
                    , Button.ripple
                    , Button.link <| "#token/" ++ toString token.id
                    , toMdlCss buttonStyle
                    , css "width" "75%"
                    , css "padding" "5px"
                    ]
                    [ text "mine"
                    ]
                ]
            ]
        ]


renderEmpty =
    div []
        [ div
            [ style
                [ ( "text-align", "center" )
                , ( "padding-top", "80px" )
                ]
            ]
            [ text "There are no tokens yet" ]
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


cardBtnsStyle =
    style
        [ ( "border-top", "1px solid #dd" )
        ]


cardBtnStyle =
    style
        [ ( "text-align", "center" )
        ]
