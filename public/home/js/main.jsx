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
                                        <table width="500px">
                                         <tr>
                                             <td>
                                                {getPlatformImage}
                                                 <img id="platformimage" src="" />
                                             </td>
                                         </tr>
                                        </table>
                                    </div>
                                    <div class="card-block">
                                        <p class="card-text">
                                            <table class="table table-vertical">
                                                <tbody>
                                                    <tr>
                                                        <th>Platform</th>
                                                        <td>{navigator.platform}</td>
                                                    </tr>
                                                    <tr>
                                                        <th>Browser Vendor</th>    
                                                        <td>{navigator.vendor}</td> 
                                                    </tr>
                                                    <tr>
                                                        <th>Browser Type</th>
                                                        <td>{navigator.product}</td>
                                                    </tr>
                                                </tbody>
                                            </table>
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

function getPlatformImage() {
    var $img = document.getElementById("platformimage")
    switch(navigator.platform) {
        case "MacIntel":
            $img.src = "/images/apple.png"
            break;
        case "Win32":
            $img.src = "/images/windows.png"
            break;
       default:
            $img.src = "/images/unicorn.png"
       break;
    }
}


setInterval(tick, 1000);