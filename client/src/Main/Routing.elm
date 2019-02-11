module Main.Routing exposing
    ( Route(..)
    , assetPath
    , exploreAssetsPath
    , fastSignupPath
    , forgotPassPath
    , homepagePath
    , loginPath
    , matchers
    , newTokenPath
    , parseLocation
    , personPath
    , portfolioPath
    , profilePath
    , signupPath
    , termsPath
    , tokenPath
    , tokensPath
    )

import Navigation exposing (Location)
import UrlParser exposing ((</>), Parser, int, map, oneOf, parseHash, s, string, top)


type Route
    = NotFoundRoute
    | HomepageRoute
    | ExploreAssetsRoute
    | TokenRoute Int
    | PersonRoute Int
    | PortfolioRoute
    | TokensRoute
    | CreateTokenRoute
    | UserLoginRoute
    | UserForgotRoute
    | UserForgotResetPassRoute Int String
    | UserSignupRoute
    | UserFastSignupRoute
    | AssetRoute Int
    | ProfileRoute Int
    | CreateAssetRoute


matchers : Parser (Route -> a) a
matchers =
    oneOf
        [ map HomepageRoute top
        , map ExploreAssetsRoute (s "explore-assets")
        , map TokenRoute (s "token" </> int)
        , map PersonRoute (s "person" </> int)
        , map PortfolioRoute (s "portfolio")
        , map TokensRoute (s "tokens")
        , map CreateTokenRoute (s "new")
        , map UserLoginRoute (s "login")
        , map UserSignupRoute (s "signup")
        , map UserFastSignupRoute (s "fsignup")
        , map UserForgotRoute (s "forgot")
        , map UserForgotResetPassRoute (s "forgot-reset-password" </> int </> string)
        , map AssetRoute (s "asset" </> int)
        , map ProfileRoute (s "profile" </> int)
        , map CreateAssetRoute (s "create-asset")
        ]


parseLocation : Location -> Route
parseLocation location =
    case parseHash matchers location of
        Just route ->
            route

        Nothing ->
            NotFoundRoute


homepagePath : String
homepagePath =
    "#"


assetPath : Int -> String
assetPath assetId =
    "#asset/" ++ toString assetId


exploreAssetsPath : String
exploreAssetsPath =
    "#explore-assets"


profilePath : Int -> String
profilePath userId =
    "#profile/" ++ toString userId


tokensPath : String
tokensPath =
    "#tokens"


tokenPath : Int -> String
tokenPath id =
    "#token/" ++ toString id


personPath : Int -> String
personPath id =
    "#person/" ++ toString id


trendingPath : String
trendingPath =
    "#tokens"


newTokenPath : String
newTokenPath =
    "#new"


portfolioPath : String
portfolioPath =
    "#portfolio"


loginPath : String
loginPath =
    "#login"


signupPath : String
signupPath =
    "#signup"


forgotPassPath : String
forgotPassPath =
    "#forgot"


fastSignupPath : String
fastSignupPath =
    "#fsignup"


termsPath : String
termsPath =
    "#site-terms"


privacyPath : String
privacyPath =
    "#site-privacy"
