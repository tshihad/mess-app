import React, { Fragment } from 'react';
import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import IconButton from '@material-ui/core/IconButton';
import Typography from '@material-ui/core/Typography';
import MenuIcon from '@material-ui/icons/Menu';
import Drawer from '@material-ui/core/Drawer';
import List from '@material-ui/core/List';
import Divider from '@material-ui/core/Divider';
import ListItem from '@material-ui/core/ListItem';
import ListItemIcon from '@material-ui/core/ListItemIcon';
import ListItemText from '@material-ui/core/ListItemText';
import Settings from '@material-ui/icons/Settings';
import PowerSettingsNew from '@material-ui/icons/PowerSettingsNew';
import AccountCircle from '@material-ui/icons/AccountCircle';

const styles = theme => ({
    root: {
        flexGrow: 1,
    },
    grow: {
        flexGrow: 1,
    },
    menuButton: {
        marginLeft: -12,
        marginRight: 20,
    },
    bottomBar: {
        top: 'auto',
        bottom: 0,
    },
    toolbar: {
        alignItems: 'center',
        justifyContent: 'space-between',
    },
});

class TopBar extends React.Component {
    state = {
        menu: false,
    };
    toggleDrawer = (side, open) => () => {
        this.setState({
            [side]: open,
        });
    };
    render() {
        const { classes } = this.props;
        const sideList = (
            <div className={classes.list}>
                <Divider />
                <List>
                    {['Settings', 'Logout'].map((text, index) => (
                        <div key={text}>
                            <ListItem button>
                                <ListItemIcon>{index % 2 === 0 ? <Settings /> : <PowerSettingsNew />}</ListItemIcon>
                                <ListItemText primary={text} />
                            </ListItem>
                            <Divider />
                        </div>
                    ))}
                </List>
            </div>
        );
        return (
            <Fragment>
                <div className={classes.root}>
                    <AppBar position="static">
                        <Toolbar>
                            <IconButton className={classes.menuButton} color="inherit" aria-label="Menu" onClick={this.toggleDrawer('menu', true)}>
                                <MenuIcon />
                            </IconButton>
                            <Typography variant="h6" color="inherit" className={classes.grow}>
                                Mess App
                        </Typography>
                        </Toolbar>
                    </AppBar>
                    <Drawer open={this.state.menu} onClose={this.toggleDrawer('menu', false)}>
                        <div
                            tabIndex={0}
                            role="button"
                            onClick={this.toggleDrawer('menu', false)}
                            onKeyDown={this.toggleDrawer('menu', false)}
                        >
                            <div className={classes.list}>
                                <Divider />
                                <List>
                                    <ListItem button>
                                        <ListItemIcon><AccountCircle /></ListItemIcon>
                                        <ListItemText primary={'Fullname'} secondary={'@username'} />
                                    </ListItem>
                                </List>
                            </div>
                        </div>
                        <Divider />
                        <div
                            tabIndex={0}
                            role="button"
                            onClick={this.toggleDrawer('menu', false)}
                            onKeyDown={this.toggleDrawer('menu', false)}
                        >
                            {sideList}
                        </div>
                    </Drawer>
                </div>
            </Fragment >
        );
    }
}

TopBar.propTypes = {
    classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(TopBar);