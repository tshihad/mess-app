import React, { Fragment } from 'react';
import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import IconButton from '@material-ui/core/IconButton';
import CssBaseline from '@material-ui/core/CssBaseline';
import HomeIcon from '@material-ui/icons/Home';
import Group from '@material-ui/icons/Group';
import Assignment from '@material-ui/icons/Assignment';
import SwapHoriz from '@material-ui/icons/SwapHoriz';
import Restaurant from '@material-ui/icons/RestaurantMenu';
import Typography from '@material-ui/core/Typography';
import MenuIcon from '@material-ui/icons/Menu';

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

function BottomAppBar(props) {
    const { classes } = props;
    return (
        <Fragment>
            <div className={classes.root}>
                <AppBar position="static">
                    <Toolbar>
                        <IconButton className={classes.menuButton} color="inherit" aria-label="Menu">
                            <MenuIcon />
                        </IconButton>
                        <Typography variant="h6" color="inherit" className={classes.grow}>
                            Mess App
                        </Typography>
                    </Toolbar>
                </AppBar>
            </div>
            <CssBaseline />
            <AppBar position="fixed" color="primary" className={classes.bottomBar}>
                <Toolbar className={classes.toolbar}>
                    <IconButton color="inherit" >
                        <HomeIcon />
                    </IconButton>
                    <IconButton color="inherit" >
                        <Group />
                    </IconButton>
                    <IconButton color="inherit" >
                        <Assignment />
                    </IconButton>
                    <IconButton color="inherit" >
                        <Restaurant />
                    </IconButton>
                    <IconButton color="inherit" >
                        <SwapHoriz />
                    </IconButton>
                </Toolbar>
            </AppBar>
        </Fragment>
    );
}

BottomAppBar.propTypes = {
    classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(BottomAppBar);