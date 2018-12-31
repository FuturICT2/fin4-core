module Actions.View exposing (render)

import Actions.Model exposing (Model)
import Actions.Msg exposing (Msg(..))
import Common.Decimal exposing (renderDecimal, renderDecimalWithPrecision)
import Common.Error as Error
import Common.Styles exposing (padding, textLeft, textRight, toMdlCss)
import Dict exposing (Dict)
import Html exposing (..)
import Html.Attributes exposing (placeholder, rows, style, type_, value)
import Html.Events exposing (onClick, onInput)
import Identicon exposing (identicon)
import Main.Context exposing (Context)
import Material.Button as Button
import Material.Chip as Chip
import Material.Icon as Icon
import Material.Options as Options exposing (css)
import Material.Table as Table
import Material.Typography as Typo
import Model.Actions exposing (Action, Actions)


render : Context -> Model -> Html Msg
render ctx model =
    let
        content =
            case model.actions of
                Just actions ->
                    case List.length actions.entries of
                        0 ->
                            renderEmpty

                        _ ->
                            renderActions ctx model actions

                Nothing ->
                    renderEmpty
    in
    content


renderActions : Context -> Model -> Actions -> Html Msg
renderActions ctx model actions =
    div
        [ style
            [ ( "margin-top", "45px" )
            ]
        ]
        [ div [] <| List.map (renderToken ctx model) actions.entries
        ]


renderToken ctx model token =
    div
        [ style
            [ ( "border-left", "10px solid #ddd" )
            , ( "margin-bottom", "15px" )
            ]
        ]
        [ renderTokenInfo ctx model token
        , renderTokenClaims ctx model token
        , renderClaimForm ctx model token

        -- , renderTokenControls ctx model token
        ]


renderTokenInfo ctx model token =
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
                [ b
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
            []
            [ div
                [ style
                    [ ( "padding", "7px" )
                    , ( "line-height", "1" )
                    , ( "font-size", "26px" )
                    ]
                ]
                [ text token.purpose
                ]
            ]
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
                    ]
                    [ Icon.view "favorite" [ Icon.size24, css "color" likeBackground ]
                    , text <| " " ++ toString token.favouriteCount
                    ]
                ]
            ]
        ]


renderTokenClaims ctx model token =
    let
        likeColor =
            case token.didUserLike of
                True ->
                    "blue"

                False ->
                    "green"

        showApproveBtn =
            case ctx.user of
                Just user ->
                    case user.id == token.creatorId of
                        True ->
                            True

                        False ->
                            False

                Nothing ->
                    False
    in
    div
        []
        [ div [ style [ ( "padding", "5px" ) ] ]
            [ case token.didUserLike of
                True ->
                    a
                        [ style [ ( "color", "blue" ), ( "padding", "15px" ) ]
                        , onClick (DoLike token.id token.didUserLike)
                        ]
                        [ text "Unlike" ]

                False ->
                    a
                        [ style [ ( "color", "green" ), ( "padding", "15px" ) ]
                        , onClick (DoLike token.id token.didUserLike)
                        ]
                        [ text "like" ]
            , text " ("
            , text <| toString token.favouriteCount
            , text " likes - "
            , text <| toString (List.length token.claims)
            , text " claims - "
            , text <|
                toString
                    (List.length <|
                        List.filter (\v -> v.isApproved == True) token.claims
                    )
            , text " approved)"
            ]
        , div
            [ style
                [ ( "text-align", "left" )
                ]
            ]
          <|
            List.map (renderClaim model showApproveBtn token.id) token.claims
        ]


renderClaim model showApproveBtn actionId claim =
    let
        btn =
            case claim.isApproved of
                True ->
                    i [] [ text " - approved" ]

                False ->
                    case showApproveBtn of
                        True ->
                            a
                                [ onClick (ApproveClaim claim.id)
                                ]
                                [ text " (approve)"
                                ]

                        False ->
                            i [] []
    in
    div
        [ style
            [ ( "padding", "10px" )
            , ( "border", "1px solid #ddd" )
            , ( "margin", "3px" )
            ]
        ]
        [ b []
            [ text <|
                claim.userName
            ]
        , btn
        , text ": "
        , text claim.text
        ]


renderClaimForm ctx model token =
    div
        [ style [ ( "margin", "10px 3px 0 3px" ) ] ]
        [ div
            []
            [ let
                v =
                    Maybe.withDefault "" <| Dict.get token.id model.claims
              in
              textarea
                [ textareaStyle
                , value v
                , placeholder "Submit proposal"
                , onInput <| SetClaim token.id
                , rows 3
                ]
                []
            , div
                [ actionButtonsStyle
                ]
                [ Button.render Mdl
                    [ token.id, -2 ]
                    model.mdl
                    [ Button.raised
                    , Button.ripple
                    , toMdlCss buttonStyle
                    , Options.onClick (SubmitProposal token.id <| Maybe.withDefault "" <| Dict.get token.id model.claims)
                    ]
                    [ text "Submit"
                    ]
                , Button.render Mdl
                    [ token.id, -1 ]
                    model.mdl
                    [ Button.raised
                    , Button.ripple
                    , toMdlCss buttonStyle
                    ]
                    [ text "Photo"
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
            [ text "No actions yet" ]
        ]


actionControlsStyle =
    style
        [ ( "margin", "15px 0" )
        , ( "border", "1px solid #ddd" )
        ]


actionControlStyle =
    style
        [ ( "border-right", "1px solid #ddd" )
        , ( "width", "33%" )
        , ( "display", "inline-block" )
        , ( "padding", "5px" )
        , ( "cursor", "pointer" )
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
        [ ( "width", "50%" )
        , ( "text-align", "center" )
        , ( "line-height", "normal" )
        , ( "text-decoration", "none" )
        , ( "color", "inherit" )
        , ( "display", "inline-block" )
        , ( "box-shadow", "none" )
        , ( "font-size", "26px" )
        , ( "padding", "15px" )
        , ( "height", "60px" )
        ]


textareaStyle : Attribute a
textareaStyle =
    style
        [ ( "background-color", "#f2f2f2" )
        , ( "border", "none" )
        , ( "padding", "4px" )
        , ( "width", "100%" )
        , ( "border-top", "1px solid #ddd" )
        , ( "box-sizing", "border-box" )
        , ( "margin-bottom", "-5px" )
        ]


cardBtnsStyle =
    style
        [ ( "border-top", "1px solid #dd" )
        ]


cardBtnStyle =
    style
        [ ( "display", "inline-block" )
        , ( "text-align", "center" )
        , ( "padding", "5px" )
        ]


imgInputStyle =
    style
        [-- ( "position", "absolute" )
         -- , ( "left", "0" )
         -- , ( "top", "0" )
         -- , ( "right", "0" )
         -- , ( "bottom", "0" )
         -- , ( "opacity", "0" )
         -- , ( "width", "100%" )
        ]
