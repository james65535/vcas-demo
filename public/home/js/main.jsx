class Clock extends React.Component {
    render() {
        return (
            <div className="main-container">
                <header className="header header-6">
                    <div _ngcontent-c0="" className="branding">
                        <a _ngcontent-c0="" className="nav-link" href="#">
                            <span _ngcontent-c0="" className="clr-icon clr-clarity-logo"></span>
                            <span _ngcontent-c0="" className="title">Clarity</span>
                        </a>
                    </div>
                    <div _ngcontent-c0="" className="header-nav clr-nav-level-1" ng-reflect-_level="1">
                        <a _ngcontent-c0="" className="nav-link" href="/home" routerlinkactive="active"
                           ng-reflect-router-link="/home" ng-reflect-router-link-active="active"><span _ngcontent-c0=""
                                                                                                       className="nav-text">Home</span></a>
                        <a _ngcontent-c0="" className="nav-link" href="/users" routerlinkactive="active"
                           ng-reflect-router-link="/users" ng-reflect-router-link-active="active"><span _ngcontent-c0=""
                                                                                                       className="nav-text">Users</span></a>
                        <a _ngcontent-c0="" className="nav-link" href="/about" routerlinkactive="active"
                           ng-reflect-router-link="/about" ng-reflect-router-link-active="active"><span _ngcontent-c0=""
                                                                                                        className="nav-text">About</span></a>
                    </div>
                    <div _ngcontent-c0="" className="header-actions">
                    </div>
                </header>
                <div className="content-container">
                    <div className="content-area">
                        <div class="row">
                            <div class="col-lg-6 col-md-12 col-sm-12 col-xs-12">
                                <a href="..." class="card clickable">
                                    <div class="card-img">
                                        <img src="/images/apple.png" />
                                    </div>
                                    <div class="card-block">
                                        <p class="card-text">
                                            Platform: {navigator.platform}
                                            <br />
                                            User Agent: {navigator.userAgent}
                                            <br />
                                            Geo: {navigator.geolocation.getCurrentPosition.name}
                                            <br />
                                            Platform: {navigator.platform}
                                            <br />
                                            Product: {navigator.product}
                                            <br />
                                            Vendor: {navigator.vendor}
                                        </p>
                                    </div>
                                </a>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        );
    }
}

function tick() {
    ReactDOM.render(
        <Clock date={new Date()} />,
        document.getElementById('root')
    );
}


setInterval(tick, 1000);