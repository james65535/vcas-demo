class Clock extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            error: null,
            isLoaded: false,
            items: []
        };
    }

    componentDidMount() {
        fetch("https://swapi.co/api/people/1/")
            .then(res => res.json())
            .then(
                (result) => {
                    this.setState({
                        isLoaded: true,
                        items: result.items
                    });
                },
                // Note: it's important to handle errors here
                // instead of a catch() block so that we don't swallow
                // exceptions from actual bugs in components.
                (error) => {
                    this.setState({
                        isLoaded: true,
                        error
                    });
                }
            )
    }
    render() {
        const { error, isLoaded, items } = this.state;
        if (error) {
            return <div>Error: {error.message}</div>;
        } else if (!isLoaded) {
            return <div>Loading...</div>;
        } else {
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
                               ng-reflect-router-link="/home" ng-reflect-router-link-active="active"><span
                                _ngcontent-c0=""
                                className="nav-text">Home</span></a>
                            <a _ngcontent-c0="" className="nav-link" href="/users" routerlinkactive="active"
                               ng-reflect-router-link="/users" ng-reflect-router-link-active="active"><span
                                _ngcontent-c0=""
                                className="nav-text">Users</span></a>
                            <a _ngcontent-c0="" className="nav-link" href="/about" routerlinkactive="active"
                               ng-reflect-router-link="/about" ng-reflect-router-link-active="active"><span
                                _ngcontent-c0=""
                                className="nav-text">About</span></a>
                        </div>
                        <div _ngcontent-c0="" className="header-actions">
                        </div>
                    </header>
                    <div className="content-container">
                        <div className="content-area">
                            <ul>
                                {items.map(item => (
                                    <li key={item.name}>
                                        {item.name} {item.height}
                                    </li>
                                ))}
                            </ul>
                        </div>
                    </div>
                </div>
            );
        }
    }
}

function tick() {
    ReactDOM.render(
        <Clock />,
        document.getElementById('root')
    );
}

setInterval(tick, 1000);