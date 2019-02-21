module Main.Update exposing (mountRoute, update)

import Asset.Command
import Asset.Model
import Asset.Update
import CreateAsset.Model
import CreateAsset.Update
import Debug
import ExploreAssets.Command
import ExploreAssets.Update
import Homepage.Command
import Homepage.Update
import Main.Command exposing (checkSessionCmd, logoutCmd)
import Main.Context exposing (DeviceSize(..))
import Main.Model exposing (Model)
import Main.Msg exposing (Msg(..))
import Main.Routing
    exposing
        ( Route(..)
        , loginPath
        , parseLocation
        )
import Material
import Navigation exposing (newUrl)
import Portfolio.Command
import Portfolio.Update
import Profile.Command
import Profile.Model
import Profile.Msg
import Profile.Update
import UserLogin.Update


mountRoute : Model -> ( Model, Cmd Msg )
mountRoute model =
    case model.context.route of
        PortfolioRoute ->
            model ! [ Cmd.map Portfolio <| Portfolio.Command.commands model.context ]

        AssetRoute _ ->
            { model | asset = Asset.Model.init }
                ! [ Cmd.map AssetMsg <| Asset.Command.commands model.context
                  ]

        ProfileRoute _ ->
            { model | profile = Profile.Model.init }
                ! [ Cmd.map ProfileMsg <| Profile.Command.commands model.context
                  ]

        HomepageRoute ->
            model
                ! [ Cmd.map HomepageMsg <| Homepage.Command.commands model.context model.homepage
                  ]

        ExploreAssetsRoute ->
            model
                ! [ Cmd.map ExploreAssetsMsg <| ExploreAssets.Command.commands model.context
                  ]

        _ ->
            model ! []


update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    let
        _ =
            Debug.log "" msg
    in
    case msg of
        OnRouteChange location ->
            let
                context =
                    model.context
            in
            mountRoute
                { model
                    | context =
                        { context
                            | route = parseLocation location
                        }
                    , showMobileNav = False
                    , createAsset = CreateAsset.Model.init
                }

        OnCheckSessionResponse resp ->
            let
                context =
                    model.context
            in
            case resp of
                Ok user ->
                    { model | context = { context | user = Just user, sessionDidLoad = True } } ! []

                Err _ ->
                    { model | context = { context | sessionDidLoad = True } } ! []

        HomepageMsg msg_ ->
            let
                ( childModel, cmd ) =
                    Homepage.Update.update model.context msg_ model.homepage
            in
            { model | homepage = childModel } ! [ Cmd.map HomepageMsg cmd ]

        ExploreAssetsMsg msg_ ->
            let
                ( childModel, cmd ) =
                    ExploreAssets.Update.update model.context msg_ model.exploreAssets
            in
            { model | exploreAssets = childModel } ! [ Cmd.map ExploreAssetsMsg cmd ]

        CreateAssetMsg msg_ ->
            let
                ( childModel, cmd ) =
                    CreateAsset.Update.update model.context msg_ model.createAsset
            in
            { model | createAsset = childModel } ! [ Cmd.map CreateAssetMsg cmd ]

        AssetMsg msg_ ->
            let
                ( childModel, cmd ) =
                    Asset.Update.update model.context msg_ model.asset
            in
            { model | asset = childModel } ! [ Cmd.map AssetMsg cmd ]

        ProfileMsg msg_ ->
            let
                loadSession =
                    case msg_ of
                        Profile.Msg.OnUploadProfileImageResponse resp ->
                            case resp of
                                Ok _ ->
                                    checkSessionCmd model.context

                                Err _ ->
                                    Cmd.none

                        _ ->
                            Cmd.none

                ( childModel, cmd ) =
                    Profile.Update.update model.context msg_ model.profile
            in
            { model | profile = childModel } ! [ Cmd.map ProfileMsg cmd, loadSession ]

        Portfolio msg_ ->
            let
                ( childModel, cmd ) =
                    Portfolio.Update.update model.context msg_ model.portfolio
            in
            { model | portfolio = childModel } ! [ Cmd.map Portfolio cmd ]

        UserLoginMsg msg_ ->
            let
                ( userlogin, userloginCmd ) =
                    UserLogin.Update.update model.context msg_ model.userlogin

                context =
                    model.context

                user =
                    case userlogin.user of
                        Just _ ->
                            userlogin.user

                        _ ->
                            context.user

                cmd =
                    case userlogin.user of
                        Just _ ->
                            if userlogin.isNewUser then
                                [ newUrl "#" ]

                            else
                                [ newUrl "#" ]

                        _ ->
                            [ Cmd.map UserLoginMsg userloginCmd ]
            in
            { model | userlogin = userlogin, context = { context | user = user } } ! cmd

        UserLogout ->
            let
                context =
                    model.context
            in
            { model | context = { context | user = Nothing } } ! [ logoutCmd context model ]

        OnLogoutResponse resp ->
            case resp of
                Ok _ ->
                    let
                        context =
                            model.context

                        cmd =
                            case context.user of
                                Just _ ->
                                    []

                                _ ->
                                    [ newUrl loginPath ]
                    in
                    { model | context = { context | user = Nothing } } ! cmd

                Err _ ->
                    model ! []

        OnWindowResize size ->
            let
                device =
                    if size.width <= 950 then
                        Mobile

                    else
                        Desktop

                window =
                    { width = size.width
                    , height = size.height
                    , device = device
                    , isMobile = device == Mobile
                    }

                context =
                    model.context
            in
            { model | context = { context | window = window } } ! []

        ToggleMobileNav ->
            { model | showMobileNav = not model.showMobileNav } ! []

        Mdl msg_ ->
            Material.update Mdl msg_ model
