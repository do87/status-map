import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { Container, Row, Col } from 'react-bootstrap';
import StatusBox from '../status-box';
import Filters from '../filters';
import './map.css';

export default () => {
    const [apiData, setApiData] = useState({});
    const [filters, setFilters] = useState({})
    const [providersProducts, setProvidersProducts] = useState({})

    useEffect(() => {
        const getProducts = async () => {
            const { data } = await axios("/api/products?groupBy=platforms")
            setApiData(data)
        };

        getProducts().catch(console.error)

        const timer = setInterval(() => {
            getProducts().catch(console.error)
        }, 10000);

        // clearing interval
        return () => clearInterval(timer);
    }, [])

    useEffect(() => {
        const initialValues = ['azlv', 'lvzt', 'live', 'silv', 'prod']
        const sortData = (data) => {
            let organizedProducts = {}
            let doInit = (filters["isInfraActive"] === undefined && data?.items !== undefined)
            let f = {
                orgs: doInit ? {} : filters.orgs || {},
                infras: doInit ? {} : filters.infras || {},
                stages: doInit ? {} : filters.stages || {},
                isInfraActive: doInit ? true : filters.isInfraActive
            }

            if (data?.items === undefined) {
                return { organizedProducts, organizedFilters: f }
            }
            data.items.forEach((provider) => {
                organizedProducts[provider.name] = provider.items
                if (doInit) {
                    provider.items.forEach((product) => {
                        f.orgs[product.org] = true
                        product.stages = product.stages || []
                        product.stages.forEach(stage => {
                            f.stages[stage] = initialValues.indexOf(stage) !== -1
                        });
                        product.infras = product.infras || []
                        product.infras.forEach(infra => {
                            f.infras[infra] = initialValues.indexOf(infra) !== -1
                        });
                    })
                } else {
                    f = filters
                }
            })
            return { organizedProducts, organizedFilters: f }
        }

        const { organizedProducts, organizedFilters } = sortData(apiData)
        setFilters(organizedFilters)
        setProvidersProducts(organizedProducts)
        // eslint-disable-next-line
    }, [apiData])

    return (
        <Container>
            <Row>
                <Col md={4}>
                    <Filters {...{ filters, setFilters }} />
                </Col>
                <Col>
                    {Object.keys(providersProducts).map(function (provider) {
                        return <Row key={`provider-${provider}`}>
                            <Col className="p-0 mt-3 mb-3 provider-row">
                                <h2>{provider.toUpperCase()}</h2>
                                <div className="status-box-wrapper p-3">
                                    {providersProducts[provider].map(product => <StatusBox key={`status-${product.org}-${product.id}`} product={product} filters={filters} />)}
                                </div>
                            </Col>
                        </Row>
                    })}
                </Col>
            </Row>
        </Container >
    );
}
