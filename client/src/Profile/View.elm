module Profile.View exposing (render)

import Common.Decimal exposing (renderDecimalWithPrecision)
import Common.Error exposing (renderHttpError)
import Common.Spinner as Spinner
import Common.Styles
    exposing
        ( marginTop
        , padding
        , paddingLeft
        , paddingRight
        , paddingTop
        , textCenter
        , textLeft
        , toMdlCss
        )
import Html exposing (..)
import Html.Attributes exposing (..)
import Html.Events exposing (on, onClick)
import Identicon exposing (identicon)
import Json.Decode as JD
import Main.Context exposing (Context)
import Main.Routing exposing (assetPath)
import Material.Button as Button
import Material.Icon as Icon
import Material.Options as Options exposing (css)
import Material.Spinner
import Material.Table as Table
import Material.Tabs as Tabs
import Material.Typography as Typography
import Profile.Model exposing (Asset, Model, Profile, UserAssets)
import Profile.Msg exposing (Msg(..), idxToTab, tabToIdx)
import Profile.Types exposing (ViewTab(..))
import Timeline.View


render : Context -> Model -> Int -> Html Msg
render ctx model profileId =
    div [ bodyStyle ]
        [ case model.data of
            Just profile ->
                renderProfile ctx model profile

            Nothing ->
                Spinner.render "One moment..."
        ]


renderProfile : Context -> Model -> Profile -> Html Msg
renderProfile ctx model profile =
    let
        imgUrl =
            case model.imgInput of
                Just data ->
                    data.contents

                Nothing ->
                    profile.profileImageUrl
    in
    div []
        [ renderHttpError model.error
        , div [ coverContainerStyle ]
            [ div [ coverInfoStyle ]
                [ div [ coverNameStyle ]
                    [ text profile.username
                    , case profile.isOwnder of
                        True ->
                            div [ style [ ( "color", "black" ) ] ]
                                [ Options.styled span
                                    [ Typography.caption ]
                                    [ text "Value of your profile is 0 FILS"
                                    ]
                                ]

                        False ->
                            span [] []
                    ]
                ]
            , div [ coverProfileImageStyle imgUrl ]
                []
            , case ctx.user of
                Just user ->
                    case user.id == profile.id of
                        True ->
                            renderImageSelector ctx model profile

                        False ->
                            span [] []

                Nothing ->
                    span [] []
            ]
        , renderTabs ctx model profile
        ]


renderTabs : Context -> Model -> Profile -> Html Msg
renderTabs ctx model profile =
    let
        opts =
            [ Options.center
            , toMdlCss (paddingLeft 18)
            , toMdlCss (paddingRight 18)
            ]

        tabs =
            case profile.isOwnder of
                True ->
                    [ Tabs.label opts
                        [ text <| "Timeline"
                        ]
                    , Tabs.label opts
                        [ text <| "Tokens"
                        ]
                    ]

                False ->
                    [ Tabs.label opts
                        [ text <| "Timeline"
                        ]
                    , Tabs.label opts
                        [ text <| "Tokens"
                        ]
                    ]
    in
    Tabs.render Mdl
        [ 0 ]
        model.mdl
        [ Tabs.ripple
        , Tabs.onSelectTab idxToTab
        , Tabs.activeTab <| tabToIdx model.activeViewTab
        , toMdlCss (paddingTop 20)
        ]
        tabs
        [ case model.activeViewTab of
            TimelineTab ->
                Html.map TimelineMsg <|
                    Timeline.View.render ctx model.timeline

            WalletTab ->
                span [] [ text "balance" ]

            TokensTab ->
                case model.userAssets of
                    Nothing ->
                        div [ paddingTop 15 ] [ text "..." ]

                    Just userAssets ->
                        renderUserAssets ctx model userAssets
        ]


renderUserAssets : Context -> Model -> UserAssets -> Html Msg
renderUserAssets ctx model asset =
    case List.length asset.entries < 1 of
        True ->
            Options.styled p
                [ Typography.body1, toMdlCss (padding 50), toMdlCss textCenter ]
                [ text "You didn't create any token yet" ]

        False ->
            div [ assetsTableStyle ]
                [ Table.table
                    []
                    [ Table.thead []
                        [ Table.tr []
                            [ Table.th [ toMdlCss textLeft ] [ text "Token" ]
                            , Table.th [ toMdlCss (paddingRight 48) ] [ text "Supply" ]
                            ]
                        ]
                    , Table.tbody [] <| List.map renderUserAsset asset.entries
                    ]
                , case asset.hasMore of
                    True ->
                        renderLoadMore ctx model

                    False ->
                        span [] []
                ]


renderUserAsset : Asset -> Html Msg
renderUserAsset asset =
    Table.tr []
        [ Table.td [ toMdlCss textLeft ]
            [ identicon "16px" asset.name
            , text " "
            , a [ href (assetPath asset.id) ] [ text asset.name ]
            ]
        , Table.td []
            [ text <| toString asset.totalSupply
            , span
                [ style
                    [ ( "display", "inline-block" )
                    , ( "text-align", "left" )
                    , ( "width", "60px" )
                    , ( "padding-left", "10px" )
                    ]
                ]
                [ text asset.symbol ]
            ]
        ]


renderLoadMore : Context -> Model -> Html Msg
renderLoadMore ctx model =
    div [ loadMoreButoonStyle ]
        [ Button.render Mdl
            [ 1 ]
            model.mdl
            [ Button.raised
            , Button.ripple
            , Button.colored
            , toMdlCss (marginTop 15)
            , Options.onClick LoadMore
            ]
            [ text "load more"
            ]
        ]


renderImageSelector : Context -> Model -> Profile -> Html Msg
renderImageSelector ctx model profile =
    div [ imgSelectorContainer ]
        [ case model.imgInput of
            Just _ ->
                div [ submitBtnStyle ]
                    [ Button.render Mdl
                        [ profile.id, -2 ]
                        model.mdl
                        [ Button.raised
                        , Button.ripple
                        , toMdlCss buttonStyle
                        , Options.onClick UploadProfileImage
                        , Options.disabled model.isUploadingImage
                        ]
                        [ case model.isUploadingImage of
                            True ->
                                Material.Spinner.spinner
                                    [ Material.Spinner.active True ]

                            False ->
                                text "save"
                        ]
                    ]

            Nothing ->
                span []
                    [ input
                        [ imgInputStyle
                        , type_ "file"
                        , id "profile-img-input"
                        , on "change"
                            (JD.succeed ImageSelected)
                        ]
                        []
                    , div [ imgIconContainerStyle ]
                        [ Icon.view "add_a_photo" [ Icon.size24 ] ]
                    ]
        ]


imgSelectorContainer : Attribute a
imgSelectorContainer =
    style
        [ ( "position", "absolute" )
        , ( "right", "0" )
        , ( "bottom", "0" )
        , ( "width", "50px" )
        ]


imgInputStyle : Attribute a
imgInputStyle =
    style
        [ ( "opacity", "0" )
        , ( "position", "absolute" )
        , ( "top", "0" )
        , ( "left", "0" )
        , ( "right", "0" )
        , ( "bottom", "0" )
        , ( "z-index", "1" )
        ]


imgIconContainerStyle : Attribute a
imgIconContainerStyle =
    style
        []


bodyStyle : Attribute a
bodyStyle =
    style
        [ ( "text-align", "center" )
        ]


coverContainerStyle : Attribute a
coverContainerStyle =
    style
        [ ( "height", "60px" )
        , ( "width", "100%" )
        , ( "background", "#42a5f5" )
        , ( "position", "relative" )
        , ( "margin-bottom", "60px" )
        , ( "color", "white" )
        ]


coverInfoStyle : Attribute a
coverInfoStyle =
    style
        [ ( "position", "absolute" )
        , ( "bottom", "10px" )
        , ( "left", "120px" )
        , ( "right", "0" )
        , ( "height", "25px" )
        , ( "line-height", "30px" )
        ]


coverProfileImageStyle : String -> Attribute a
coverProfileImageStyle data =
    let
        url =
            case data == "" of
                True ->
                    "images/default-user-icon.jpg"

                False ->
                    data
    in
    style
        [ ( "position", "absolute" )
        , ( "bottom", "-50px" )
        , ( "left", "10px" )
        , ( "width", "100px" )
        , ( "height", "100px" )
        , ( "border", "3px solid white" )
        , ( "border-radius", "50%" )
        , ( "background", "url(" ++ url ++ ") no-repeat center /cover" )
        ]


coverNameStyle : Attribute a
coverNameStyle =
    style
        [ ( "display", "inline-block" )
        , ( "float", "left" )
        , ( "font-weight", "bold" )
        , ( "font-size", "18px" )
        , ( "text-align", "left" )
        ]


submitBtnStyle : Attribute a
submitBtnStyle =
    style
        [ ( "text-align", "center" )
        ]


buttonStyle : Attribute a
buttonStyle =
    style
        [ ( "text-align", "center" )
        , ( "line-height", "normal" )
        , ( "text-decoration", "none" )
        , ( "color", "inherit" )
        , ( "display", "inline-block" )
        , ( "box-shadow", "none" )
        , ( "float", "right" )
        , ( "border-radius", "8px" )
        , ( "background", "white" )
        , ( "color", "black" )
        , ( "margin", "5px" )
        ]


assetsTableStyle : Attribute a
assetsTableStyle =
    style
        [ ( "font-size", "12px" )
        , ( "margin", "10px" )
        ]


loadMoreButoonStyle : Attribute a
loadMoreButoonStyle =
    style [ ( "text-align", "center" ) ]
