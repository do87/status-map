import React from 'react';
import { Row } from 'react-bootstrap';
import Checkbox from '@mui/material/Checkbox';
import Radio from '@mui/material/Radio';
import RadioGroup from '@mui/material/RadioGroup';
import FormControlLabel from '@mui/material/FormControlLabel';
import FormGroup from '@mui/material/FormGroup';

import { createTheme, ThemeProvider } from '@mui/material/styles';
import './filters.css';

const darkTheme = createTheme({
    palette: {
        mode: 'dark',
    },
});

export default ({ filters, setFilters }) => {


    const setIsInfraActive = (flag = true) => {
        setFilters((fs) => {
            let fc = { ...fs }
            fc.isInfraActive = flag
            return fc
        })
    }

    const setOrg = (org, ischecked) => {
        setFilters((fs) => {
            let fc = { ...fs }
            fc.orgs[org] = ischecked.checked
            return fc
        })
    }

    const setInfra = (infra, ischecked) => {
        setFilters((fs) => {
            let fc = { ...fs }
            fc.infras[infra] = ischecked.checked
            return fc
        })
    }

    const setStage = (stage, ischecked) => {
        setFilters((fs) => {
            let fc = { ...fs }
            fc.stages[stage] = ischecked.checked
            return fc
        })
    }

    if (!Object.keys(filters).length) {
        return <></>
    }

    return <div className="p-2 filters">
        <h2><i className="fa fa-filter"></i> Filter</h2>
        <Row>
            <ThemeProvider theme={darkTheme}>

                {Object.keys(filters.orgs).sort().map(org => {
                    return <FormControlLabel
                        control={<Checkbox sx={{
                            color: '#fff',
                            '&.Mui-checked': {
                                color: '#fff',
                            },
                        }} checked={filters.orgs[org]} />}
                        key={`org-cbox-${org}`}
                        id={`org-${org}`}
                        label={org}
                        onClick={(e) => setOrg(org, e.target)}
                    />
                })}
                <hr />

                <RadioGroup
                    value={filters.isInfraActive || filters.isInfraActive === undefined ? "infra" : "stage"}
                    name="radio-buttons-group"
                    row
                >
                    <FormControlLabel
                        control={<Radio sx={{
                            color: '#fff',
                            '&.Mui-checked': {
                                color: '#fff',
                            },
                        }} />}
                        label="Infra"
                        value="infra"
                        onClick={setIsInfraActive}
                    />

                    <FormControlLabel
                        control={<Radio sx={{
                            color: '#fff',
                            '&.Mui-checked': {
                                color: '#fff',
                            },
                        }} />}
                        label="Stage"
                        value="stage"
                        onClick={() => setIsInfraActive(false)}
                    />
                </RadioGroup>
                <hr />

                <FormGroup row>
                    {filters.isInfraActive ? Object.keys(filters.infras).sort().map(infra => {
                        return <FormControlLabel
                            style={{ width: "40%" }}
                            control={<Checkbox sx={{
                                color: '#fff',
                                '&.Mui-checked': {
                                    color: '#fff',
                                },
                            }} checked={filters.infras[infra]} />}
                            key={`infra-cbox-${infra}`}
                            id={`infra-${infra}`}
                            label={infra}

                            onClick={(e) => setInfra(infra, e.target)}
                        />
                    }) : Object.keys(filters.stages).sort().map(stage => {
                        return <FormControlLabel
                            control={<Checkbox sx={{
                                color: '#fff',
                                '&.Mui-checked': {
                                    color: '#fff',
                                },
                            }} checked={filters.stages[stage]} />}
                            style={{ width: "40%" }}
                            key={`stage-cbox-${stage}`}
                            id={`stage-${stage}`}
                            label={stage}
                            onClick={(e) => setStage(stage, e.target)}
                        />
                    })}
                </FormGroup>
            </ThemeProvider>
        </Row>
    </div >
}
