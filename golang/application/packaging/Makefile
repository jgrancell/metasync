package/binaries:
	bash packaging/generate-binaries.sh ${BINARY} ${WORKDIR}

package/containers: binaries
	bash packaging/generate-containers.sh ${WORKDIR}/packaging

package/clean:
	rm -rf packaging/artifacts/${BINARY}-*
	rm -rf packaging/binaries/${BINARY}-*