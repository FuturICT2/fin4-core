module Token.View exposing (render)

import Common.Decimal exposing (renderDecimalWithPrecision)
import Common.Error exposing (renderHttpError)
import Common.Spinner as Spinner
import Common.Styles exposing (toMdlCss)
import Html exposing (..)
import Html.Attributes exposing (..)
import Html.Events exposing (on, onClick, onInput)
import Identicon exposing (identicon)
import Json.Decode as JD
import Main.Context exposing (Context)
import Main.Routing exposing (personPath)
import Material
import Material.Button as Button
import Material.Icon as Icon
import Material.Options as Options exposing (css)
import Material.Tabs as Tabs
import Model.Tokens exposing (Claim, Liker, Miner, Token)
import Token.Model exposing (Model)
import Token.Msg exposing (Msg(..))


render : Context -> Model -> Int -> Html Msg
render ctx model tokenId =
    div [ style [ ( "text-align", "center" ), ( "margin-bottom", "60px" ) ] ]
        [ case model.data of
            Nothing ->
                renderSpinner

            Just token ->
                renderToken ctx model token
        ]


renderSpinner : Html Msg
renderSpinner =
    div []
        [ Spinner.render "One moment..." ]


renderToken : Context -> Model -> Token -> Html Msg
renderToken ctx model token =
    div []
        [ div
            [ style
                [ ( "padding", "20px" )
                , ( "text-align", "center" )
                , ( "width", "80px" )
                , ( "height", "80px" )
                , ( "border", "1px solid #ddd" )
                , ( "border-radius", "50%" )
                , ( "margin", "15px auto 0 auto" )
                ]
            ]
            [ identicon "40px" token.name
            ]
        , header [] [ text token.name ]
        , div [ style [ ( "margin-bottom", "15px" ) ] ]
            [ text token.purpose
            ]
        , case model.showClaimForm of
            True ->
                div []
                    [ renderClaimForm ctx model token
                    , Button.render Mdl
                        [ 0 ]
                        model.mdl
                        [ Button.ripple
                        , Options.onClick ToggleClaimButton
                        ]
                        [ text "cancel" ]
                    ]

            False ->
                div []
                    [ renderToggleClaimButton ctx model token
                    , renderTabs ctx model token
                    ]
        ]


renderTabs : Context -> Model -> Token -> Html Msg
renderTabs ctx model token =
    Tabs.render Mdl
        [ 0 ]
        model.mdl
        [ Tabs.ripple
        , Tabs.onSelectTab SelectTab
        , Tabs.activeTab model.tab
        ]
        [ Tabs.label
            [ Options.center ]
            [ text <| "Claims " ++ toString (List.length token.claims)
            ]
        , Tabs.label
            [ Options.center ]
            [ text <| "Miners " ++ toString (List.length token.miners)
            ]
        , Tabs.label
            [ Options.center ]
            [ text <| "Likers " ++ toString token.favouriteCount
            ]
        ]
        [ case model.tab of
            0 ->
                renderTokenClaims ctx model token

            1 ->
                renderMiners ctx model token

            _ ->
                renderLikers ctx model token
        ]


renderLikers : Context -> Model -> Token -> Html Msg
renderLikers ctx model token =
    case List.length token.likers of
        0 ->
            div
                [ style
                    [ ( "padding", "15px" )
                    ]
                ]
                [ text "There are no likers yet"
                ]

        _ ->
            div [] <|
                List.map (renderLiker token) token.likers


renderLiker : Token -> Liker -> Html Msg
renderLiker token liker =
    div
        [ style
            [ ( "text-align", "left" )
            , ( "border", "1px solid #ddd" )
            , ( "border-radius", "8px" )
            , ( "margin", "5px 0" )
            , ( "padding", "5px" )
            ]
        ]
        [ a [] [ text liker.userName ]
        ]


renderMiners : Context -> Model -> Token -> Html Msg
renderMiners ctx model token =
    case List.length token.miners of
        0 ->
            div
                [ style
                    [ ( "padding", "15px" )
                    ]
                ]
                [ text "There are no miners yet"
                ]

        _ ->
            div [] <|
                List.map (renderMiner token) token.miners


renderMiner : Token -> Miner -> Html Msg
renderMiner token miner =
    div
        [ style
            [ ( "text-align", "left" )
            , ( "border", "1px solid #ddd" )
            , ( "border-radius", "8px" )
            , ( "margin", "5px 0" )
            , ( "padding", "5px" )
            ]
        ]
        [ a [] [ text miner.userName ]
        , div
            [ style
                [ ( "float", "right" )
                ]
            ]
            [ text "Mined: "
            , renderDecimalWithPrecision miner.mined 2
            , text <| " " ++ token.symbol
            ]
        ]


renderTokenClaims : Context -> Model -> Token -> Html Msg
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
    case List.length token.claims of
        0 ->
            div
                [ style
                    [ ( "padding", "15px" )
                    ]
                ]
                [ text "There are no claims yet"
                ]

        _ ->
            div
                [ style
                    [ ( "text-align", "left" )
                    ]
                ]
            <|
                List.map (renderClaim model showApproveBtn token.id) token.claims


renderClaim : Model -> Bool -> Int -> Claim -> Html Msg
renderClaim model showApproveBtn tokenId claim =
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
            , ( "border-radius", "8px" )
            , ( "margin", "10px 0" )
            ]
        ]
        [ a [ href <| personPath claim.userId ]
            [ text <|
                claim.userName
            ]
        , btn
        , text ": "
        , text claim.text
        , case claim.logoFile == "" of
            True ->
                span [] []

            False ->
                div
                    [ style
                        [ ( "width", "100px" )
                        , ( "height", "100px" )
                        , ( "margin-bottom", "15px" )
                        , ( "background", "url(" ++ claim.logoFile ++ ") no-repeat center center" )
                        , ( "background-size", "contain" )
                        ]
                    ]
                    []
        ]


renderToggleClaimButton : Context -> Model -> Token -> Html Msg
renderToggleClaimButton ctx model token =
    Button.render Mdl
        [ 0 ]
        model.mdl
        [ Button.raised
        , Button.ripple
        , Options.onClick ToggleClaimButton
        , toMdlCss claimButtonStyle
        ]
        [ text <| "Mine " ++ token.symbol ]


renderClaimForm : Context -> Model -> Token -> Html Msg
renderClaimForm ctx model token =
    let
        v =
            Maybe.withDefault { contents = "", filename = "" } model.img
    in
    div
        [ claimFormContainerStyle ]
        [ div
            []
            [ renderHttpError model.submitProposalError
            , div
                [ imgSelectorContainer v.contents
                ]
                [ input
                    [ imgInputStyle
                    , type_ "file"
                    , id "claim-img-proposal"
                    , on "change"
                        (JD.succeed ImageSelected)
                    ]
                    []
                , div [ imgIconContainerStyle ] [ Icon.view "add_a_photo" [ Icon.size48 ] ]
                ]
            , case v.contents == "" of
                True ->
                    div [] []

                False ->
                    div []
                        [ img
                            [ src v.contents
                            , style
                                [ ( "max-width", "90%" )
                                , ( "max-height", "200px" )
                                , ( "margin-bottom", "25px" )
                                ]
                            ]
                            []
                        ]
            , textarea
                [ textareaStyle
                , value model.proposalText
                , placeholder "claim text"
                , onInput SetClaim
                , rows 3
                ]
                []

            ----------------------------------------------------
            , div [ cardBtnsStyle ]
                [ div
                    [ cardBtnStyle
                    ]
                    [ Button.render Mdl
                        [ token.id, -2 ]
                        model.mdl
                        [ Button.raised
                        , Button.ripple
                        , toMdlCss buttonStyle
                        , Options.onClick (SubmitProposal model.proposalText)
                        ]
                        [ text "Submit"
                        ]
                    ]
                ]

            ----------------------------------------------------
            ]
        ]


imgInputStyle : Attribute a
imgInputStyle =
    style
        [ ( "position", "absolute" )
        , ( "left", "0" )
        , ( "top", "0" )
        , ( "right", "0" )
        , ( "bottom", "0" )
        , ( "opacity", "0" )
        , ( "width", "100%" )
        , ( "z-index", "1" )
        , ( "padding-top", "70px" )
        ]


tokenControlsStyle : Attribute a
tokenControlsStyle =
    style
        [ ( "margin", "15px 0" )
        , ( "border", "1px solid #ddd" )
        ]


tokenControlStyle : Attribute a
tokenControlStyle =
    style
        [ ( "border-right", "1px solid #ddd" )
        , ( "width", "33%" )
        , ( "display", "inline-block" )
        , ( "cursor", "pointer" )
        ]


tokenButtonsStyle : Attribute a
tokenButtonsStyle =
    style
        [ ( "border-top", "1px solid #ddd" )
        , ( "background", "#f2f2f2" )
        , ( "font-weight", "bold" )
        , ( "display", "block" )
        ]


buttonStyle : Attribute a
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


cardBtnsStyle : Attribute a
cardBtnsStyle =
    style
        [ ( "border-top", "1px solid #dd" )
        ]


cardBtnStyle : Attribute a
cardBtnStyle =
    style
        [ ( "text-align", "center" )
        ]


claimButtonStyle : Attribute a
claimButtonStyle =
    style
        [ ( "width", "100%" )
        , ( "margin", "15px 0" )
        ]


claimFormContainerStyle : Attribute a
claimFormContainerStyle =
    style
        [ ( "margin", "10px 0px 0 0px" )
        , ( "border", "1px solid #ddd" )
        , ( "border-radius", "8px" )
        , ( "margin-bottom", "30px" )
        ]


imgSelectorContainer : String -> Attribute a
imgSelectorContainer bgContents =
    style
        [ -- ( "background-image", "URL(" ++ bgContents ++ ") no-repeat center center fixed" )
          -- , ( "background-size", "contain" )
          -- , ( "background-repeat", "no-repeat" )
          ( "position", "relative" )
        , ( "width", "100%" )
        ]


imgIconContainerStyle : Attribute a
imgIconContainerStyle =
    style
        [ ( "margin-top", "15px" )
        ]
