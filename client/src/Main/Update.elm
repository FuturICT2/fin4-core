module Main.Update exposing (mountRoute, update)

import Actions.Command
import Actions.Model
import Actions.Update
import CreateAction.Model
import CreateAction.Update
import CreateToken.Model
import CreateToken.Update
import Debug
import Main.Command exposing (logoutCmd)
import Main.Context exposing (DeviceSize(..))
import Main.Model exposing (Model)
import Main.Msg exposing (Msg(..))
import Main.Routing
    exposing
        ( Route(..)
        , loginPath
        , parseLocation
        , tokensPath
        )
import Material
import Navigation exposing (newUrl)
import Portfolio.Command
import Portfolio.Model
import Portfolio.Update
import Token.Command
import Token.Model
import Token.Update
import Tokens.Command
import Tokens.Update
import UserLogin.Update


mountRoute : Model -> ( Model, Cmd Msg )
mountRoute model =
    case model.context.route of
        TokensRoute ->
            model ! [ Cmd.map Tokens <| Tokens.Command.commands model.context ]

        TokenRoute tokenId ->
            let
                token =
                    Token.Model.init
            in
            { model | token = token } ! [ Cmd.map Token <| Token.Command.commands model.context tokenId ]

        PortfolioRoute ->
            model ! [ Cmd.map Portfolio <| Portfolio.Command.commands model.context ]

        ActionsRoute ->
            model ! [ Cmd.map Actions <| Actions.Command.commands model.context ]

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
                    , createToken = CreateToken.Model.init
                    , createAction = CreateAction.Model.init
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

        Homepage msg_ ->
            model ! []

        Tokens msg_ ->
            let
                ( childModel, cmd ) =
                    Tokens.Update.update model.context msg_ model.tokens
            in
            { model | tokens = childModel } ! [ Cmd.map Tokens cmd ]

        Token msg_ ->
            let
                ( childModel, cmd ) =
                    Token.Update.update model.context msg_ model.token
            in
            { model | token = childModel } ! [ Cmd.map Token cmd ]

        Portfolio msg_ ->
            let
                ( childModel, cmd ) =
                    Portfolio.Update.update model.context msg_ model.portfolio
            in
            { model | portfolio = childModel } ! [ Cmd.map Portfolio cmd ]

        Actions msg_ ->
            let
                ( childModel, cmd ) =
                    Actions.Update.update model.context msg_ model.actions
            in
            { model | actions = childModel } ! [ Cmd.map Actions cmd ]

        CreateToken msg_ ->
            let
                ( childModel, cmd ) =
                    CreateToken.Update.update model.context msg_ model.createToken
            in
            { model | createToken = childModel } ! [ Cmd.map CreateToken cmd ]

        CreateAction msg_ ->
            let
                ( childModel, cmd ) =
                    CreateAction.Update.update model.context msg_ model.createAction
            in
            { model | createAction = childModel } ! [ Cmd.map CreateAction cmd ]

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
                                [ newUrl "#actions" ]

                            else
                                [ newUrl "#actions" ]

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
