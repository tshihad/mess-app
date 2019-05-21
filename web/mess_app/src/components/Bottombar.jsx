import React, { Fragment } from 'react';
import PropTypes from 'prop-types';
import { withStyles } from '@material-ui/core/styles';
import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import IconButton from '@material-ui/core/IconButton';
import CssBaseline from '@material-ui/core/CssBaseline';
import HomeIcon from '@material-ui/icons/Home';
import Group from '@material-ui/icons/Group';
import Receipt from '@material-ui/icons/Receipt';
import SwapHoriz from '@material-ui/icons/SwapHoriz';
import Restaurant from '@material-ui/icons/RestaurantMenu';

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
                        <Receipt />
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