import React, { useState, useEffect } from 'react';
import './status-box.css';

export default ({ product, filters }) => {
    const [computedStatus, setStatus] = useState("");
    const [computedProduct, setComputedProduct] = useState({});

    useEffect(() => {
        if (filters.isInfraActive) {
            setComputedProduct(() => {
                let newProd = JSON.parse(JSON.stringify(product))
                newProd.infras = newProd.infras || []

                newProd.infras.forEach(infra => {
                    if (!filters.infras[infra]) {
                        delete newProd.status_infras[infra]
                        return
                    }
                })
                newProd.infras = Object.keys(newProd.status_infras)
                return newProd
            })
        } else {
            setComputedProduct(() => {
                // console.log(product)
                let newProd = JSON.parse(JSON.stringify(product))
                newProd.stages = newProd.stages || []

                newProd.stages.forEach(stage => {
                    // console.log(stage, filters.stages[stage])
                    if (!filters.stages[stage]) {
                        delete newProd.status_stages[stage]
                        return
                    }
                })
                newProd.stages = Object.keys(newProd.status_stages)
                // console.log(newProd)
                return newProd
            })
        }
    }, [product, filters])

    useEffect(() => {
        if (!Object.keys(computedProduct).length) {
            return
        }

        let statusMap = {
            'error': 0,
            'running': 0,
            'done': 0,
            'unknown': 0
        }

        const analyzeStatusMap = (activeFlag, map) => {
            if (!map?.current_status || !activeFlag) {
                return
            }

            switch (map.current_status) {
                case 'done':
                    statusMap['done']++
                    break;
                case 'running':
                    statusMap['running']++
                    break;
                case 'error':
                    statusMap['error']++
                    break;
                default:
                    statusMap['unknown']++
                    break;
            }
        }

        if (filters.isInfraActive) {
            Object.keys(filters.infras).forEach(infra => {
                analyzeStatusMap(filters.infras[infra], computedProduct.status_infras[infra])
            })
        } else {
            Object.keys(filters.stages).forEach(stage => {
                analyzeStatusMap(filters.stages[stage], computedProduct.status_stages[stage])
            })
        }

        if (Object.values(statusMap).reduce((a, b) => a + b) === 0) {
            return
        }

        Object.keys(statusMap).forEach((s) => {
            if (statusMap[s] > 0) {
                setStatus(s)
                return
            }
        })
    }, [filters, computedProduct])

    if (computedStatus === "") {
        return <></>
    }

    if (!Object.keys(filters).length) {
        return <></>
    }

    if (filters.isInfraActive && computedProduct.infras?.length === 0) {
        return <></>
    }

    if (!filters.isInfraActive && (computedProduct.stages?.length === 0)) {
        return <></>
    }

    if (!filters.orgs[computedProduct.org]) {
        return <></>
    }

    return <div className={"status-box-label status-" + computedStatus}>
        <a href={`https://odj.cloud/prods/show?product=${computedProduct.id}`} target="_blank" rel="noopener noreferrer">
            <div className="status-box-label-text">
                {computedProduct.id}
            </div>
        </a>
    </div>
}

// error, done, running