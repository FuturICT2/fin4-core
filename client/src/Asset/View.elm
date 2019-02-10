module Asset.View exposing (render)

import Asset.Model exposing (Asset, Block, Image, Miner, Model)
import Asset.Msg exposing (Msg(..))
import Common.Decimal exposing (renderDecimalWithPrecision)
import Common.Error exposing (renderHttpError)
import Common.Spinner as Spinner
import Common.Styles
    exposing
        ( background
        , padding
        , paddingBottom
        , paddingLeft
        , paddingRight
        , paddingTop
        , textCenter
        , textLeft
        , toMdlCss
        )
import Html exposing (..)
import Html.Attributes exposing (..)
import Html.Events exposing (on, onClick, onFocus, onInput)
import Identicon exposing (identicon)
import Json.Decode as JD
import Main.Context exposing (Context)
import Main.Routing exposing (profilePath)
import Material.Button as Button
import Material.Elevation as Elevation
import Material.Icon as Icon
import Material.Options as Options exposing (css)
import Material.Spinner
import Material.Table as Table
import Material.Tabs as Tabs
import Material.Typography as Typography
import Timeline.View


render : Context -> Model -> Int -> Html Msg
render ctx model assetId =
    div []
        [ case model.data of
            Just asset ->
                renderAsset ctx model asset

            Nothing ->
                Spinner.render "One moment..."
        ]


renderAsset : Context -> Model -> Asset -> Html Msg
renderAsset ctx model asset =
    div [ assetCardStyle ctx.window.isMobile ]
        [ div [ assetCoverContainerStyle ]
            [ div [ assetCoverInfoStyle ]
                [ div [ coverNameStyle ]
                    [ text asset.name
                    , div
                        [ style
                            [ ( "margin-top", "5px" )
                            , ( "color", "black" )
                            ]
                        ]
                        [ Options.styled span
                            [ Typography.caption ]
                            [ text <|
                                "Supply= "
                                    ++ toString asset.totalSupply
                                    ++ " "
                                    ++ asset.symbol
                            , text " | Market Cap= 0FILS"
                            ]
                        ]
                    , div [ style [ ( "color", "black" ) ] ]
                        [ Options.styled span
                            [ Typography.caption ]
                            [ text "Moderator: "
                            , a
                                [ href (profilePath asset.creatorId) ]
                                [ text asset.creatorName ]
                            ]
                        ]
                    ]
                ]
            , div [ assetQRStyle ]
                [ identicon "60px" asset.name
                ]
            ]
        , div [ assetPurposeStyle ]
            [ Options.styled p
                [ Typography.body1 ]
                [ text asset.description ]
            ]
        , renderTabs ctx model asset
        ]


renderTabs : Context -> Model -> Asset -> Html Msg
renderTabs ctx model asset =
    let
        opts =
            [ Options.center
            , toMdlCss (paddingLeft 18)
            , toMdlCss (paddingRight 18)
            ]
    in
    Tabs.render Mdl
        [ 0 ]
        model.mdl
        [ Tabs.ripple
        , Tabs.onSelectTab SelectTab
        , Tabs.activeTab model.tab
        ]
        [ Tabs.label opts
            [ text <| "Timeline "
            ]
        , Tabs.label opts
            [ text <| "Miners " ++ toString (List.length asset.miners)
            ]
        , Tabs.label opts
            [ text <| "Trade"
            ]
        ]
        [ case model.tab of
            0 ->
                renderBlocksTab ctx model asset

            1 ->
                renderMiners ctx model asset

            _ ->
                renderTrade ctx model asset
        ]


renderBlocksTab : Context -> Model -> Asset -> Html Msg
renderBlocksTab ctx model asset =
    div []
        [ renderBlockForm ctx model asset
        , Options.styled p
            [ Typography.caption, toMdlCss textLeft, toMdlCss (paddingLeft 10) ]
            [ text <| "When the moderator ("
            , a [ href (profilePath asset.creatorId) ]
                [ text asset.creatorName ]
            , text <|
                ") accepts your post 1"
                    ++ asset.symbol
                    ++ " will be mined to your balance"
            ]
        , renderBlocksList ctx model asset
        ]


renderBlocksList : Context -> Model -> Asset -> Html Msg
renderBlocksList ctx model asset =
    Html.map TimelineMsg <|
        Timeline.View.render ctx model.timeline


renderBlockForm : Context -> Model -> Asset -> Html Msg
renderBlockForm ctx model asset =
    Options.div
        [ Elevation.e0
        , toMdlCss blockFormContainerStyle
        ]
        [ div
            []
            [ textarea
                [ textareaStyle
                , value model.blockText
                , placeholder <| "Compose new post! "
                , onInput SetBlockText
                , rows 5
                ]
                []
            , renderFormImages ctx model asset
            , renderHttpError model.submitBlockError
            , div [ cardBtnsStyle ]
                [ renderImageSelector ctx model asset
                , div [ submitBtnStyle ]
                    [ Button.render Mdl
                        [ asset.id, -2 ]
                        model.mdl
                        [ Button.raised
                        , Button.ripple
                        , Button.colored
                        , toMdlCss buttonStyle
                        , Options.onClick SubmitBlock
                        , Options.disabled model.isSubmittingBlock
                        ]
                        [ case model.isSubmittingBlock of
                            True ->
                                Material.Spinner.spinner
                                    [ Material.Spinner.active True ]

                            False ->
                                text "Publish"
                        ]
                    ]
                ]
            ]
        ]


renderImageSelector : Context -> Model -> Asset -> Html Msg
renderImageSelector ctx model asset =
    case List.length model.imgs > 2 of
        True ->
            small [ imgSelectorContainer ] [ text "Images limit reached" ]

        False ->
            div [ imgSelectorContainer ]
                [ input
                    [ imgInputStyle
                    , type_ "file"
                    , id "block-img"
                    , on "change"
                        (JD.succeed ImageSelected)
                    ]
                    []
                , div [ imgIconContainerStyle ]
                    [ Icon.view "add_a_photo" [ Icon.size24 ] ]
                ]


renderFormImages : Context -> Model -> Asset -> Html Msg
renderFormImages ctx model asset =
    div [ imgsContainerStyle ] <|
        List.indexedMap (renderImage ctx model asset) model.imgs


renderImage : Context -> Model -> Asset -> Int -> Image -> Html Msg
renderImage context model asset idx image =
    case image.contents == "" of
        True ->
            div [] []

        False ->
            div [ imgThumbnailContainerStyle ]
                [ img [ src image.contents, imgThumbnailStyle ] []
                , div [ imgDeleteStyle, onClick (DeleteImage idx) ] [ text "X" ]
                ]


renderMiners : Context -> Model -> Asset -> Html Msg
renderMiners context model asset =
    case List.length asset.miners < 1 of
        True ->
            div [ paddingTop 15 ] [ text "No mining happened yet" ]

        False ->
            div [ minersTableStyle ]
                [ Table.table
                    []
                    [ Table.thead []
                        [ Table.tr []
                            [ Table.th [] [ text "User" ]
                            , Table.th []
                                [ text <| "Mined (" ++ asset.symbol ++ ")" ]
                            , Table.th [] [ text "Mining Percentage" ]
                            ]
                        ]
                    , Table.tbody [] <|
                        List.map (renderMiner context model asset) asset.miners
                    ]
                ]


renderMiner : Context -> Model -> Asset -> Miner -> Html Msg
renderMiner context model asset miner =
    Table.tr []
        [ Table.td []
            [ a [ href <| profilePath miner.id ] [ text miner.userName ]
            ]
        , Table.td [] [ renderDecimalWithPrecision miner.mined 2 ]
        , Table.td []
            [ renderDecimalWithPrecision miner.miningPercentage 2
            , text "%"
            ]
        ]


renderTrade : Context -> Model -> Asset -> Html Msg
renderTrade context model asset =
    div [ padding 50, textCenter ]
        [ Options.styled p
            [ Typography.caption ]
            [ text "Coming soon..." ]
        ]


assetCardStyle : Bool -> Attribute a
assetCardStyle isMobile =
    let
        minWidth =
            case isMobile of
                True ->
                    "300px"

                False ->
                    "600px"
    in
    style
        [ ( "padding-bottom", "50px" )
        , ( "text-align", "center" )
        , ( "max-width", "600px" )
        , ( "min-width", minWidth )
        , ( "margin", "auto" )
        ]



---------------------------------------------------


assetCoverContainerStyle : Attribute a
assetCoverContainerStyle =
    style
        [ ( "height", "60px" )
        , ( "width", "100%" )
        , ( "background", "#42a5f5" )
        , ( "position", "relative" )
        , ( "margin-bottom", "60px" )
        , ( "color", "white" )
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


assetCoverInfoStyle : Attribute a
assetCoverInfoStyle =
    style
        [ ( "position", "absolute" )
        , ( "bottom", "0" )
        , ( "left", "120px" )
        , ( "right", "0" )
        , ( "height", "25px" )
        , ( "li ne-height", "40px" )
        ]


assetQRStyle : Attribute a
assetQRStyle =
    style
        [ ( "position", "absolute" )
        , ( "bottom", "-50px" )
        , ( "left", "10px" )
        , ( "width", "100px" )
        , ( "height", "100px" )
        , ( "border", "3px solid white" )
        , ( "border-radius", "50%" )
        , ( "background", "white" )
        , ( "padding-top", "20px" )
        ]



---------------------------------------------------


identiconStyle : Attribute a
identiconStyle =
    style
        [ ( "width", "35px" )
        , ( "height", "35px" )
        , ( "padding", "5px" )
        , ( "border", "1px solid #ddd" )
        , ( "background", "white" )
        , ( "border-radius", "50%" )
        , ( "margin", "0 10px" )
        , ( "display", "inline-block" )
        ]


buttonStyle : Attribute a
buttonStyle =
    style
        [ ( "text-align", "center" )
        , ( "float", "right" )
        ]


cardBtnsStyle : Attribute a
cardBtnsStyle =
    style
        [ ( "border-top", "1px solid #dd" )
        , ( "height", "45px" )
        ]


submitBtnStyle : Attribute a
submitBtnStyle =
    style
        [ ( "text-align", "center" )
        ]


imgSelectorContainer : Attribute a
imgSelectorContainer =
    style
        [ ( "position", "relative" )
        , ( "float", "left" )
        , ( "padding", "12px 0 0 3px" )
        ]


imgIconContainerStyle : Attribute a
imgIconContainerStyle =
    style
        []


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
        ]


textareaStyle : Attribute a
textareaStyle =
    style
        [ ( "border", "none" )
        , ( "padding", "4px 8px" )
        , ( "width", "100%" )
        , ( "box-sizing", "border-box" )
        , ( "margin-top", "5px" )
        ]


blockFormContainerStyle : Attribute a
blockFormContainerStyle =
    style
        [ ( "margin", "10px" )
        , ( "border", "1px solid #ddd" )
        , ( "padding", "0 5px" )
        ]


imgsContainerStyle : Attribute a
imgsContainerStyle =
    style
        [ ( "text-align", "left" )
        ]


imgThumbnailContainerStyle : Attribute a
imgThumbnailContainerStyle =
    style
        [ ( "text-align", "left" )
        , ( "position", "relative" )
        , ( "display", "inline-block" )
        , ( "margin", "15px 20px 10px 0" )
        ]


imgThumbnailStyle : Attribute a
imgThumbnailStyle =
    style
        [ ( "width", "80px" )
        , ( "height", "80px" )
        ]


imgDeleteStyle : Attribute a
imgDeleteStyle =
    style
        [ ( "width", "25px" )
        , ( "height", "25px" )
        , ( "position", "absolute" )
        , ( "top", "-10px" )
        , ( "right", "-10px" )
        , ( "border-radius", "50%" )
        , ( "background", "red" )
        , ( "padding", "4px 7px" )
        , ( "color", "white" )
        , ( "opacity", ".8" )
        ]


blocksListStyle : Attribute a
blocksListStyle =
    style [ ( "text-align", "left" ) ]


minersTableStyle : Attribute a
minersTableStyle =
    style
        [ ( "font-size", "12px" )
        , ( "margin", "10px" )
        ]


assetPurposeStyle : Attribute a
assetPurposeStyle =
    style
        [ ( "margin", "15px 10px" )
        ]


change24Style : Attribute a
change24Style =
    style
        [ ( "float", "right" )
        , ( "display", "inline-block" )
        , ( "margin-right", "10px" )
        , ( "padding-top", "5px" )
        ]
