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
                        <div>
                            <h1>Hello, yall</h1>
                            <h2>It is {this.props.date.toLocaleTimeString()}.</h2>
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