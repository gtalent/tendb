
import {
	BrowserRouter,
	Link,
	Route,
} from 'react-router-dom';
import React from 'react';
import AppBar from '@material-ui/core/AppBar';
import Button from '@material-ui/core/Button';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import Registration from './Registration';
import './App.css';

const App: React.FC = () => {
	return (
		<BrowserRouter>
			<AppBar style={{flexGrow: 1}} position="static">
				<Toolbar>
					<Typography variant="h6" color="inherit" style={{flexGrow: 1}}>10db</Typography>
					<Button color="inherit" component={props => <Link to='/register/' {...props} />}>Register</Button>
					<Button color="inherit" component={props => <Link to='/signin/' {...props} />}>Sign In</Button>
				</Toolbar>
			</AppBar>
			<Route path="/register/" render={() => ( <Registration/> )}/>
			<Route path="/signin/" render={() => ( <div>Sign In Page</div> )}/>
		</BrowserRouter>
	);
}

export default App;
