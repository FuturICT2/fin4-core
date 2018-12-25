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
        [ div [] <| List.map (renderRow ctx model) actions.entries
        ]


renderRow ctx model action =
    let
        showApproveBtn =
            case ctx.user of
                Just user ->
                    case user.id == action.creatorId of
                        True ->
                            True

                        False ->
                            False

                Nothing ->
                    False

        ( status, statusBG, statusColor ) =
            case action.status of
                0 ->
                    ( "new", "green", "white" )

                1 ->
                    ( "unconfirmed", "orange", "white" )

                2 ->
                    ( "done", "grey", "black" )

                _ ->
                    ( "unkown", "none", "black" )

        endsIn =
            case action.endsInHours of
                "1" ->
                    "remaining time: " ++ action.endsInMinutes ++ "mins"

                "0" ->
                    "remaining time: " ++ action.endsInMinutes ++ "mins"

                "-0" ->
                    ""

                _ ->
                    "ends in " ++ action.endsInHours ++ "hrs"
    in
    div
        [ style
            [ ( "border", "1px solid #ddd" )
            , ( "margin-bottom", "30px" )
            , ( "border-radius", "8px" )
            ]
        ]
        [ -- Chip.span
          --     [ css "background" statusBG
          --     , css "color" statusColor
          --     ]
          --     [ Chip.content [] [ text status ] ]
          -- ,
          div
            [ style
                [ ( "padding", "5px 0 0 5px" )
                ]
            ]
            [ div
                [ style
                    [ ( "padding", "8px" )
                    ]
                ]
                [ b [] [ text <| action.creatorName ++ ": " ]
                , text action.description
                ]
            ]
        , div
            []
            [ let
                v =
                    Maybe.withDefault "" <| Dict.get action.id model.proposals
              in
              textarea
                [ textareaStyle
                , value v
                , placeholder "Submit proposal"
                , onInput <| SetProposal action.id
                , rows 3
                ]
                []
            , div
                [ actionButtonsStyle
                ]
                [ Button.render Mdl
                    [ action.id ]
                    model.mdl
                    [ Button.raised
                    , Button.ripple
                    , toMdlCss buttonStyle
                    , Options.onClick (SubmitProposal action.id <| Maybe.withDefault "" <| Dict.get action.id model.proposals)
                    ]
                    [ text "Submit"

                    -- , text <| " " ++ toString token.favouriteCount
                    ]
                ]

            -- , div []
            --     [ input
            --         [ imgInputStyle
            --         , type_ "file"
            --
            --         -- , id model.imageId
            --         -- , on "change"
            --         --     (JD.succeed ImageSelected)
            --         ]
            --         []
            --     ]
            ]
        , div [ style [ ( "padding", "5px" ) ] ]
            [ text <| toString (List.length action.proposals)
            , text " proposals"
            ]
        , div
            [ style
                [ ( "border-top", "1px solid #ddd" )
                , ( "border-bottom", "1px solid #ddd" )
                , ( "text-align", "left" )
                ]
            ]
          <|
            List.map (renderProposal model showApproveBtn action.id) action.proposals
        ]


renderReward reward =
    span []
        [ text reward.userName
        , text "+"
        , renderDecimalWithPrecision reward.amount 2
        , text " "
        ]


renderProposal model showApproveBtn actionId proposal =
    let
        btn =
            case showApproveBtn of
                True ->
                    a
                        [ onClick (ApproveProposal proposal.id actionId proposal.doerId)
                        ]
                        [ text "Approve"
                        ]

                False ->
                    span [] []
    in
    div
        [ style
            [ ( "padding", "10px" )
            , ( "border-bottom", "1px solid #ddd" )
            ]
        ]
        [ b []
            [ text <|
                proposal.doerName
            ]
        , text " ("
        , btn
        , text "): "
        , text proposal.description
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
        [ ( "width", "100%" )
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
        , ( "padding", "15px" )
        , ( "width", "100%" )
        , ( "border-top", "1px solid #ddd" )
        , ( "box-sizing", "border-box" )
        , ( "margin-bottom", "-5px" )
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
